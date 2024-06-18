package sparrow

import (
	"github.com/lrayt/small-sparrow/application"
	"github.com/lrayt/small-sparrow/application/common"
	"github.com/lrayt/small-sparrow/kit/config_manager"
	"github.com/lrayt/small-sparrow/kit/log_manager"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var app *application.App

func Init(name, version string, options ...application.Option) error {
	if app != nil {
		return nil
	}
	app = application.NewApp(name, version)
	for _, option := range options {
		option(app)
	}

	// default workdir
	if len(app.Env.WorkDir) <= 0 {
		exePath, err := os.Executable()
		if err != nil {
			return err
		}
		application.WithWorkerDir(filepath.Dir(exePath))(app)
	}

	// print
	app.Print()

	// default configurator
	if app.ConfigProvider == nil {
		if provider, err := config_manager.NewYamlConfigProvider(app.Env); err != nil {
			return err
		} else {
			application.WithConfigurator(provider)(app)
		}
	}

	// default logger
	if app.LoggerProvider == nil {
		if provider, err := log_manager.NewLocalFileLogProvider(app.Env); err != nil {
			return err
		} else {
			application.WithLogger(provider)(app)
		}
	}
	return nil
}

func AddHandler(providers ...common.HandleProvider) {
	application.WithHandler(providers)(app)
}

func Setup() {
	var (
		errChan    = make(chan error, 1)
		signalChan = make(chan os.Signal, 1)
	)

	for _, provider := range app.HandleProviders {
		if provider == nil {
			continue
		}
		go func(fn common.HandleProvider) {
			if err := fn.Run(); err != nil {
				errChan <- err
			}
		}(provider)
	}

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errChan:
		log.Fatalf("服务启动异常，err:%v", err)
	case <-signalChan:
		for _, provider := range app.HandleProviders {
			if err := provider.Shutdown(); err != nil {
				log.Printf("shutdown server:%s\n", err.Error())
			}
		}
	}
}

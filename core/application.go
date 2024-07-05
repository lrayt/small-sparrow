package core

import (
	"fmt"
	"github.com/lrayt/small-sparrow/core/abstract"
	"github.com/lrayt/small-sparrow/core/runtime"
	"github.com/lrayt/small-sparrow/kit/config_manager"
	"github.com/lrayt/small-sparrow/kit/log_manager"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type Application struct {
	Env            *runtime.Env
	LicenseChecker abstract.LicenseChecker // 证书校验器
	LoggerProvider abstract.LoggerProvider
	ConfigProvider abstract.ConfigProvider
	Starters       []abstract.Starter
	Handlers       []abstract.Handler
}

func (app Application) Print() {
	fmt.Printf("AppName: %s\n", app.Env.AppName)
	fmt.Printf("RunEnv: %s\n", app.Env.RunEnv)
	fmt.Printf("Version: %s\n", app.Env.BuildVersion)
	fmt.Printf("WorkDir: %s\n", app.Env.WorkDir)
}

type Option func(app *Application)

func WithLicenseChecker(checker abstract.LicenseChecker) Option {
	return func(app *Application) {
		app.LicenseChecker = checker
	}
}

func WithWorkerDir(dir string) Option {
	return func(app *Application) {
		app.Env.WorkDir = dir
	}
}

func WithLogger(provider abstract.LoggerProvider) Option {
	return func(app *Application) {
		app.LoggerProvider = provider
	}
}

func WithConfigurator(provider abstract.ConfigProvider) Option {
	return func(app *Application) {
		app.ConfigProvider = provider
	}
}

func WithStarter(starters ...abstract.Starter) Option {
	return func(app *Application) {
		app.Starters = starters
	}
}

func WithHandler(handlers ...abstract.Handler) Option {
	return func(app *Application) {
		app.Handlers = handlers
	}
}

var app = new(Application)

func InitApp(appName, version string, options ...Option) error {
	app.Env = runtime.NewEnv(appName, version)
	for _, option := range options {
		option(app)
	}
	// default workdir
	if len(app.Env.WorkDir) <= 0 {
		if err := app.Env.SetDefaultWorkDir(); err != nil {
			return err
		}
	}
	// print
	app.Print()
	// license verify
	if app.LicenseChecker != nil {
		if err := app.LicenseChecker.Verify(); err != nil {
			return err
		}
	}
	// default configurator
	if app.ConfigProvider == nil {
		if provider, err := config_manager.NewYamlConfigProvider(app.Env); err != nil {
			return err
		} else {
			WithConfigurator(provider)(app)
		}
	}
	// default logger
	if app.LoggerProvider == nil {
		if provider, err := log_manager.NewLocalFileLogProvider(app.Env); err != nil {
			return err
		} else {
			WithLogger(provider)(app)
		}
	}
	return nil
}

func SetupApp() {
	var (
		errChan    = make(chan error, 1)
		signalChan = make(chan os.Signal, 1)
	)
	for _, starter := range app.Starters {
		if err := starter.Init(); err != nil {
			log.Fatalf("启动失败，err:%s\n", err.Error())
		} else {
			log.Printf("%s初始化成功\n", reflect.TypeOf(starter).String())
		}
	}
	for _, provider := range app.Handlers {
		if provider == nil {
			continue
		}
		go func(fn abstract.Handler) {
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
		// shutdown handler
		for _, handler := range app.Handlers {
			if err := handler.Shutdown(); err != nil {
				log.Printf("shutdown handler:%s\n", err.Error())
			}
		}
		// close starter
		for _, starter := range app.Starters {
			if err := starter.Close(); err != nil {
				log.Printf("%s close err: %s\n", reflect.TypeOf(starter).String(), err.Error())
			} else {
				log.Printf("%s closed\n", reflect.TypeOf(starter).String())
			}
		}
	}
}

package application

import (
	"fmt"
	"github.com/lrayt/small-sparrow/application/common"
	"github.com/lrayt/small-sparrow/application/runtime"
)

type App struct {
	Env             *runtime.Env
	LoggerProvider  common.LoggerProvider
	ConfigProvider  common.ConfigProvider
	HandleProviders []common.HandleProvider
}

func NewApp(name, version string) *App {
	appEnv := &runtime.Env{AppName: name, BuildVersion: version}
	appEnv.LoadRunEnv()
	return &App{Env: appEnv}
}

func (app App) Print() {
	fmt.Printf("AppName: %s\n", app.Env.AppName)
	fmt.Printf("RunEnv: %s\n", app.Env.RunEnv)
	fmt.Printf("Version: %s\n", app.Env.BuildVersion)
	fmt.Printf("WorkDir: %s\n", app.Env.WorkDir)
}

type Option func(app *App)

func WithWorkerDir(dir string) Option {
	return func(app *App) {
		app.Env.WorkDir = dir
	}
}

func WithLogger(provider common.LoggerProvider) Option {
	return func(app *App) {
		app.LoggerProvider = provider
	}
}

func WithConfigurator(provider common.ConfigProvider) Option {
	return func(app *App) {
		app.ConfigProvider = provider
	}
}

func WithHandler(providers []common.HandleProvider) Option {
	return func(app *App) {
		app.HandleProviders = providers
	}
}

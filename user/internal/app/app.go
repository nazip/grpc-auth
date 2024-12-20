package app

import (
	"context"
	"github.com/nazip/grpc-auth/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	//configFile      string
	//httpServer      desc.ServerHttpChi
	//metricsServer desc.ServerMetrics
}

func NewApp(ctx context.Context) (*App, error) {
	a := new(App)

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		//a.initServiceProvider,
		//a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	return config.Load(".env")
}

func (a *App) Run() error {

	return nil
}

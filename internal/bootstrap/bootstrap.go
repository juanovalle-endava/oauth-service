package bootstrap

import (
	"context"
	"fmt"
	"github.com/juanovalle-endava/oauth-service/internal/api"
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/controller"
	"github.com/juanovalle-endava/oauth-service/internal/token"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	token.Module,
	controller.Module,
	api.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle, routes api.Api,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("--------------- Starting OAuth Service ---------------")

			go func() {
				routes.SetupRoutes()
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("--------------- Stopping OAuth Servuce ---------------")
			return nil
		},
	})
}

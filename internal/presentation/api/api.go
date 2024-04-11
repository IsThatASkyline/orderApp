package api

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/routes"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/engine"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/di/api"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	router engine.RequestHandler,
	routers routes.Routes, //nolint:all
) {
	routers.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("Recovered when boot api server, r %s", r)
						}
					}()
					err := router.Fiber.Listen((":8080"))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping api application")
				return nil
			},
		},
	)
}

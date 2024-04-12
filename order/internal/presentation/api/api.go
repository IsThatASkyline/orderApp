package api

import (
	"context"
	"fmt"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/logger"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/config"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/controllers/routes"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/engine"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/api"
	fibRec "github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	router engine.RequestHandler,
	logger logger.Logger,
	config config.APIConfig,
	routers routes.Routes, //nolint:all
) {
	routers.Setup()
	router.Fiber.Use(fibRec.New())
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info(fmt.Sprintf("Starting application in :%d", config.Port))
				go func() {
					defer func() {
						if r := recover(); r != nil {
							logger.Info(fmt.Sprintf("Recovered when boot api server, r %s", r))
						}
					}()
					err := router.Fiber.Listen(fmt.Sprintf("%s:%d", config.Host, config.Port))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping api application")
				return nil
			},
		},
	)
}

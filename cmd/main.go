package main

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/di"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		api.Module,
	).Run()
}

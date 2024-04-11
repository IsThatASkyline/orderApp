package engine

import (
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/engine"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	engine.NewRequestHandler,
)

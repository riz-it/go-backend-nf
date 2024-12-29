//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"riz.it/nurul-faizah/internal/config"
	"riz.it/nurul-faizah/internal/delivery/controller"
	"riz.it/nurul-faizah/internal/delivery/middleware"
	"riz.it/nurul-faizah/internal/delivery/route"
	"riz.it/nurul-faizah/internal/domain"
	"riz.it/nurul-faizah/internal/repository"
	"riz.it/nurul-faizah/internal/usecase"
)

var authSet = wire.NewSet(
	repository.NewUserAccount,
	wire.Bind(new(domain.UserAccountRepository), new(*repository.UserAccountRepository)),
	usecase.NewAuthUseCase,
	controller.NewAuthController,
)

var classSet = wire.NewSet(
	repository.NewClass,
	wire.Bind(new(domain.ClassRepository), new(*repository.ClassRepository)),
	usecase.NewClassUseCase,
	controller.NewClassController,
)
var middlewareSet = wire.NewSet(
	middleware.NewAuthMiddleware,
)

func InitializedApp() *config.App {
	wire.Build(
		config.Get,
		config.NewLogger,
		config.NewDatabase,
		config.NewValidator,
		config.NewFiber,
		config.NewApp,
		route.NewRouter,
		usecase.NewJWTHelperImpl,
		authSet,
		classSet,
		middlewareSet,
	)
	return nil
}

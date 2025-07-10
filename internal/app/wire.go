//go:build wireinject
// +build wireinject

package app

import (
	"database/sql"
	"go-machine-boilerplate/internal/user/adapter"
	"go-machine-boilerplate/internal/user/service"
	"go-machine-boilerplate/internal/user/transport/web"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// user domain set
	adapter.NewUserAdapterPostgres,
	wire.Bind(new(service.UserRepository), new(*adapter.UserAdapterPostgres)),
	service.NewUserService,
)

func InitializeUserRoute(r *chi.Mux, db *sql.DB) (web.Route, error) {
	panic(wire.Build(
		ProviderSet,
		web.NewHandler,
		wire.Bind(new(web.Route), new(*web.Handler)),
	))
}

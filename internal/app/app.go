package app

import (
	"database/sql"
	"fmt"
	"go-machine-boilerplate/pkg/utils/httpserver"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func Run() error {
	// load env config
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
		),
	)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	// init http router and middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// add user web routes
	_, err = InitializeUserRoute(r, db)
	if err != nil {
		return err
	}

	fmt.Println("running server at port ", config.ServiceAddress)
	httpserver.Serve(config.ServiceAddress, "tcp", r)

	return nil
}

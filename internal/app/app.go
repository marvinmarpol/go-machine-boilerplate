package app

import (
	"fmt"
	"go-machine-boilerplate/pkg/utils/httpserver"
)

func Run() error {

	config, err := loadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	httpserver.Serve(config.ServiceAddress, "tcp", nil)

	return nil
}

package app

import (
	"fmt"
	"go-machine-boilerplate/internal/accessmanager/service"
)

func Run() error {

	accessManagerService := service.NewAccessManagerService()
	accessManagerService.AddRole("admin")
	accessManagerService.AddResource("kubernetes")

	accessManagerService.AddPermissions("admin", "kubernetes", []string{"read"})
	accessManagerService.AddPermissions("admin", "kubernetes", []string{"write"})

	fmt.Println(accessManagerService.CheckAccess("admin", "kubernetes", "write"))

	return nil
}

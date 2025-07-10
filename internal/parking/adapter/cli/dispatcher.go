package cli

import (
	"fmt"
	"go-machine-boilerplate/internal/parking/service"
)

const (
	createParkingCommand = "create_parking_lot"
	parkVehicleCommand   = "park_vehicle"
	unparkVehicleCommand = "unpark_vehicle"
	displayCommand       = "display"
	unknownCommand       = "\nunknown_command\n\n"

	commandErrorF = "\ninvalid arguments for %s\n\n"
)

type Command struct {
	Name string
	Args []string
}

func (cmd *Command) Dispatch(parkingLotService *service.ParkingLotService) {
	switch cmd.Name {
	case createParkingCommand:
		if len(cmd.Args) != 3 {
			fmt.Printf(commandErrorF, createParkingCommand)
			break
		}
		parkingLotService.CreateParkingLot(cmd.Args[0], cmd.Args[1], cmd.Args[2])

	case parkVehicleCommand:
		if len(cmd.Args) != 3 {
			fmt.Printf(commandErrorF, parkVehicleCommand)
			break
		}
		parkingLotService.ParkVehicle(cmd.Args[0], cmd.Args[1], cmd.Args[2])

	case unparkVehicleCommand:
		if len(cmd.Args) != 1 {
			fmt.Printf(commandErrorF, unparkVehicleCommand)
			break
		}
		parkingLotService.UnparkVehicle(cmd.Args[0])

	case displayCommand:
		if len(cmd.Args) != 2 {
			fmt.Printf(commandErrorF, displayCommand)
			break
		}
		parkingLotService.Display(cmd.Args[0], cmd.Args[1])

	default:
		fmt.Print(unknownCommand)
	}
}

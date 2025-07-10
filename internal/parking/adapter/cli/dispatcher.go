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
	unknownCommand       = "unknown_command"

	commandErrorF = "invalid arguments for %s"
)

type Command struct {
	Name string
	Args []string
}

func (cmd *Command) Dispatch(parkingLotService *service.ParkingLotService) {
	switch cmd.Name {
	case createParkingCommand:
		if len(cmd.Args) != 3 {
			fmt.Println(fmt.Sprintf(commandErrorF, createParkingCommand))
			break
		}
		parkingLotService.CreateParkingLot(cmd.Args[0], cmd.Args[1], cmd.Args[2])

	case parkVehicleCommand:
		if len(cmd.Args) != 3 {
			fmt.Println(fmt.Sprintf(commandErrorF, parkVehicleCommand))
			break
		}

	case unparkVehicleCommand:
		if len(cmd.Args) != 1 {
			fmt.Println(fmt.Sprintf(commandErrorF, unparkVehicleCommand))
			break
		}

	case displayCommand:
		if len(cmd.Args) != 1 {
			fmt.Println(fmt.Sprintf(commandErrorF, displayCommand))
			break
		}

	default:
		fmt.Println(unknownCommand)
	}
}

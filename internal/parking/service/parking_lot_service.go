package service

import (
	"fmt"
	"go-machine-boilerplate/internal/parking/domain"
	"strconv"
	"strings"
)

type ParkingLotService struct {
	parkingLot *domain.ParkingLot
}

func NewParkingLotService() *ParkingLotService {
	return &ParkingLotService{}
}

func (s *ParkingLotService) CreateParkingLot(id, floorStr, SlotStr string) {

	floors, _ := strconv.Atoi(floorStr)
	slots, _ := strconv.Atoi(SlotStr)

	if floors < 1 || slots < 1 {
		fmt.Printf("\tFailed to create parking lot. Invalid format with %d floors and %d slots per floor\n", floors, slots)
		return
	}

	parkingLot := &domain.ParkingLot{
		ID: id,
	}

	for f := 1; f <= floors; f++ {
		floor := domain.NewFloor(f)
		for s := 1; s <= slots; s++ {
			slot := domain.NewSlot(s)
			floor.Slots = append(floor.Slots, slot)
			floor.Info.MapSlotType[slot.Type]++
			floor.Info.MapOccupiedSlot[slot.Type][slot.Number] = false
		}

		parkingLot.Floors = append(parkingLot.Floors, floor)
	}

	fmt.Printf("\tCreated parking lot %s with %d floors and %d slots per floor\n", id, floors, slots)
	s.parkingLot = parkingLot
}

func (s *ParkingLotService) ParkVehicle(vStypeStr, regNumber, color string) {
	vType := domain.VehicleType(vStypeStr)

	if s.parkingLot == nil {
		fmt.Printf("\tParking lot not created yet\n")
		return
	}

	vehicle := &domain.Vehicle{
		RegistrationNumber: regNumber,
		Color:              domain.VehicleColor(strings.ToUpper(color)),
		Type:               vType,
	}

	for _, floor := range s.parkingLot.Floors {
		for _, slot := range floor.Slots {
			if !slot.Occupied && slot.Type == vType {
				slot.Occupied = true
				slot.Vehicle = vehicle
				floor.Info.MapOccupiedSlot[slot.Type][slot.Number] = true
				ticket := s.parkingLot.GenerateTicket(s.parkingLot.ID, floor.Number, slot)
				fmt.Printf("\tParked vehicle. Ticket ID: %s\n", ticket.ID)
				return
			}
		}
	}

	fmt.Printf("\tParking Lot Full\n")
}

func (s *ParkingLotService) UnparkVehicle(ticket string) {
	parts := strings.SplitN(ticket, "_", 3)
	if len(parts) != 3 {
		fmt.Printf("\tInvalid Ticket\n")
		return
	}

	floorNumber, _ := strconv.Atoi(parts[1])
	slotNumber, _ := strconv.Atoi(parts[2])

	if floorNumber < 1 || slotNumber < 1 {
		fmt.Printf("\tInvalid Ticket\n")
		return
	}

	slot := s.parkingLot.Floors[floorNumber-1].Slots[slotNumber-1]
	if slot.Occupied {
		fmt.Printf("\tUnparked vehicle with Registration Number: %s and Color: %s\n", slot.Vehicle.RegistrationNumber, slot.Vehicle.Color)
		slot.Occupied = false
		slot.Vehicle = nil
		s.parkingLot.Floors[floorNumber-1].Info.MapOccupiedSlot[slot.Type][slotNumber] = false
	} else {
		fmt.Println("Invalid Ticket")
	}

}

func (s *ParkingLotService) Display(displayType, vStypeStr string) {
	vType := domain.VehicleType(vStypeStr)

	if s.parkingLot == nil {
		fmt.Printf("\tParking lot not created yet\n")
		return
	}

	switch displayType {
	case "free_count":
		for _, floor := range s.parkingLot.Floors {
			fmt.Printf("\tNo. of slots for %s on floor %d: %d\n", vType, floor.Number, floor.GetSlotCountByType(vType))
		}
	case "occupied_slots":
		for _, floor := range s.parkingLot.Floors {
			fmt.Printf("\tOccupied slots for %s on floor %d: %s\n", vType, floor.Number, strings.Join(floor.GetSlotByOccupationStatus(vType, true), ", "))
		}
	case "free_slots":
		for _, floor := range s.parkingLot.Floors {
			fmt.Printf("\tFree slots for %s on floor %d: %s\n", vType, floor.Number, strings.Join(floor.GetSlotByOccupationStatus(vType, false), ", "))
		}
	default:
		fmt.Printf("\tInvalid display type\n")
	}
}

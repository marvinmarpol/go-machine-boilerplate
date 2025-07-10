package service

import (
	"fmt"
	"go-machine-boilerplate/internal/parking/domain"
	"strconv"
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

	parkingLot := &domain.ParkingLot{
		ID: id,
	}

	for f := 1; f < floors; f++ {
		floor := &domain.Floor{Number: f}
		for s := 1; s < slots; s++ {
			floor.Slots = append(floor.Slots, &domain.Slot{Number: s})
		}

		parkingLot.Floors = append(parkingLot.Floors, floor)
	}

	fmt.Printf("Created parking lot with %d floors and %d slots per floor", floors, slots)
	s.parkingLot = parkingLot
}

func (s *ParkingLotService) GetParkingLot() *domain.ParkingLot {
	return s.parkingLot
}

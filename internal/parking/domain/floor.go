package domain

import "fmt"

type Floor struct {
	Number int
	Slots  []*Slot
	Info   FloorInformation
}

type FloorInformation struct {
	MapSlotType     map[VehicleType]int
	MapOccupiedSlot map[VehicleType]map[int]bool
}

func NewFloor(number int) *Floor {
	mapOccupiedSlot := map[VehicleType]map[int]bool{}
	for _, v := range mapVehicleType {
		mapOccupiedSlot[v] = map[int]bool{}
	}

	return &Floor{
		Number: number,
		Info: FloorInformation{
			MapSlotType:     make(map[VehicleType]int),
			MapOccupiedSlot: mapOccupiedSlot,
		},
	}
}

func (f *Floor) GetSlotCountByType(vType VehicleType) int {
	return f.Info.MapSlotType[vType]
}

func (f *Floor) GetSlotByOccupationStatus(vType VehicleType, occupied bool) []string {
	result := []string{}

	for k, v := range f.Info.MapOccupiedSlot[vType] {
		if v == occupied {
			result = append(result, fmt.Sprintf("%d", k))
		}
	}

	return result
}

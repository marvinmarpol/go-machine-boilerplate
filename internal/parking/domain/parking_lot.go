package domain

import "fmt"

type ParkingLot struct {
	ID     string
	Floors []*Floor
}

type Ticket struct {
	ID          string
	FloorNumber int
	Slot        *Slot
}

func (p *ParkingLot) GenerateTicket(id string, floorNumber int, slot *Slot) Ticket {
	return Ticket{
		ID:          fmt.Sprintf("%s_%d_%d", id, floorNumber, slot.Number),
		Slot:        slot,
		FloorNumber: floorNumber,
	}
}

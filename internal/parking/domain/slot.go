package domain

type Slot struct {
	Number   int
	Type     VehicleType
	Occupied bool
	Vehicle  *Vehicle
}

func NewSlot(number int) *Slot {
	return &Slot{
		Number:   number,
		Type:     GetVehicleType(number),
		Occupied: false,
		Vehicle:  nil,
	}
}

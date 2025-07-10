package domain

type VehicleType string
type VehicleColor string

const (
	Car   VehicleType = "CAR"
	Bike  VehicleType = "BIKE"
	Truck VehicleType = "TRUCK"

	White VehicleColor = "WHITE"
	Black VehicleColor = "BLACK"
	Red   VehicleColor = "RED"
)

var (
	mapVehicleType = map[int]VehicleType{
		1: Truck,
		2: Bike,
	}

	mapVehicleColor = map[VehicleColor]bool{
		White: true,
		Black: true,
		Red:   true,
	}
)

type Vehicle struct {
	RegistrationNumber string
	Color              VehicleColor
	Type               VehicleType
}

func GetVehicleType(key int) VehicleType {
	value, ok := mapVehicleType[key]
	if ok {
		return value
	}

	return Car
}

package service

import (
	"go-machine-boilerplate/pkg/utils/stdhelper"
	"strings"
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	service := NewParkingLotService()

	output := stdhelper.CaptureOutput(func() {
		service.CreateParkingLot("PR1234", "2", "6")
		service.Display("free_count", "CAR")
		service.Display("free_count", "BIKE")
		service.Display("free_count", "TRUCK")
	})

	expected := []string{
		"Created parking lot PR1234 with 2 floors and 6 slots per floor",
		"No. of slots for CAR on floor 1: 3",
		"No. of slots for CAR on floor 2: 3",
		"No. of slots for BIKE on floor 1: 2",
		"No. of slots for BIKE on floor 2: 2",
		"No. of slots for TRUCK on floor 1: 1",
		"No. of slots for TRUCK on floor 2: 1",
	}

	for _, substring := range expected {
		if !strings.Contains(output, substring) {
			t.Errorf("Expected output to contain:  %s\ngot: %s", substring, output)
		}
	}

}

func TestParkAndUnpart(t *testing.T) {
	service := NewParkingLotService()

	truckOutputList := strings.Split(
		stdhelper.CaptureOutput(func() {
			service.CreateParkingLot("ABCD1234", "1", "4")
			service.ParkVehicle("TRUCK", "RG1234NM", "RED")
		}),
		"\n",
	)

	expectedSubstrings := []string{
		"Created parking lot ABCD1234 with 1 floors and 4 slots per floor",
		"Parked vehicle. Ticket ID: ABCD1234_1_1",
	}

	for i, expected := range expectedSubstrings {
		if strings.TrimSpace(truckOutputList[i]) != expected {
			t.Errorf("Expected output to contain: '%s'\ngot: '%s'", expected, truckOutputList[i])
		}
	}

	truckUnpark := strings.Split(
		stdhelper.CaptureOutput(func() {
			service.UnparkVehicle("ABCD1234_1_1")
		}),
		"\n",
	)

	expectedSubstring := "Unparked vehicle with Registration Number: RG1234NM and Color: RED"

	if strings.TrimSpace(truckUnpark[0]) != expectedSubstring {
		t.Errorf("Expected out to contain `%s`\n got `%s`", expectedSubstring, truckUnpark[0])
	}

	failedUnpark := strings.Split(
		stdhelper.CaptureOutput(func() {
			service.UnparkVehicle("RANDOM_TIKET_1234")
		}),
		"\n",
	)
	expectedSubstring = "Invalid Ticket"

	if strings.TrimSpace(failedUnpark[0]) != expectedSubstring {
		t.Errorf("Expected out to contain `%s`\n got `%s`", expectedSubstring, failedUnpark[0])
	}

}

func TestParkingFull(t *testing.T) {
	service := NewParkingLotService()

	output := strings.Split(
		stdhelper.CaptureOutput(func() {
			service.CreateParkingLot("ABCD1234", "1", "1")
			service.ParkVehicle("CAR", "B-222-AA", "WHITE")
		}),
		"\n",
	)

	expectedSubstrings := []string{
		"Created parking lot ABCD1234 with 1 floors and 1 slots per floor",
		"Parking Lot Full",
	}

	for i, expected := range expectedSubstrings {
		if strings.TrimSpace(output[i]) != expected {

			t.Errorf("Expected output %s\ngot %s", expected, output[i])
		}
	}

}

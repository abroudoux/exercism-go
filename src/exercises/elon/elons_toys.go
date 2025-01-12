package elon

import "fmt"

func (car *Car) Drive() {
	if car.batteryDrain > car.battery {
		return
	}

	car.battery = car.battery - car.batteryDrain
	car.distance = car.distance + car.speed
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
    totalDistance := (car.battery / car.batteryDrain) * car.speed
    return totalDistance >= trackDistance
}
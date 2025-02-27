package elon

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

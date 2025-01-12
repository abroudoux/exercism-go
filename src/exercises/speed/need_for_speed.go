package speed

type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery: 100,
		batteryDrain: batteryDrain,
		speed: speed,
		distance: 0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}

	return Car{
		battery: car.battery - car.batteryDrain,
		batteryDrain: car.batteryDrain,
		speed: car.speed,
		distance: car.distance + car.speed,
	}
}

func CanFinish(car Car, track Track) bool {
	return car.battery >= car.batteryDrain * (track.distance / car.speed)
}
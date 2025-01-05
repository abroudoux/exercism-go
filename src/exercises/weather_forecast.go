// Package weather provides methods for weather reporting.
// in the exercise, the package is "weather"
package exercises

// CurrentCondition is the local weather condition.
var CurrentCondition string
// CurrentLocation is the current location.
var CurrentLocation string

// Forecast returns a string with the current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
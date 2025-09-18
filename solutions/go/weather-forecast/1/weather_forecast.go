// Package weather reports weather forecasts if different cities.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation is the name of the location whose weather is being reported.
var CurrentLocation string

// Forecast returns the weather forecast at the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

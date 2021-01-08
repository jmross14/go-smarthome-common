package gosmarthomecommon

// TemperatureReading contains a temperature reading for a DHT-22 sensor.
type TemperatureReading struct {
	Humidity float64 `json:"humudity"`
	Temperature float64 `json:"temperature"`
}
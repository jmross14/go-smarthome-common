//Package gosmarthomecommon contains common items used across multiple go-smarthome projects.
package gosmarthomecommon

// TemperatureReading contains a temperature reading for a DHT-22 sensor.
type TemperatureReading struct {
	Humidity float64 `json:"humudity"`
	Temperature float64 `json:"temperature"`
}

// Status is used to signify a status
type Status int

const (
	// Off signifies the Off status
	Off Status = iota
	// On signifies the On status
	On Status = iota
)

// TriggerRelay is used to trigger a relay on/off
type TriggerRelay struct {
	Status Status
}

// SensorType is meant to distinguish the type of sensor.
type SensorType int

const (
	// TemperatureSensor Distinguishes that a sensor is a temperature sensor.
	TemperatureSensor SensorType = iota
)

// SensorModel is meant to distinguish the model of an Item.
type SensorModel string

const (
	// DHT22 distinguishes a DHT-22 Humidity/Temperature sensor
	DHT22 SensorModel = "DHT-22"
)

// SensorConnectMessage is sent from a sensor to a hub so that the hub can make periodic calls to the sensor for readings.
type SensorConnectMessage struct {
	SensorType SensorType
	Model SensorModel
	IP string
	Port string
}
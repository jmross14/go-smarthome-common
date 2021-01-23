package structs

type action func()

// DeviceType is meant to distinguish the type of sensor.
type DeviceType int

const (
	// TemperatureSensor Distinguishes that a device is a temperature sensor.
	TemperatureSensor DeviceType = iota
	// GarageOpener Distinguishes that device is a garage door opener.
	GarageOpener
	
)

// SensorModel is meant to distinguish the model of an Item.
type SensorModel string

const (
	// DHT22 distinguishes a DHT-22 Humidity/Temperature sensor
	DHT22 SensorModel = "DHT-22"
)

// SensorConnectMessage is sent from a sensor to a hub so that the hub can make periodic calls to the sensor for readings.
type SensorConnectMessage struct {
	SensorType DeviceType
	Model SensorModel
	IP string
	Port string
}

// MessageType defines an enum that contains the type of the message.
type MessageType int

const (
	// ConnectionMessage is used when connecting to another system.
	ConnectionMessage MessageType = iota
	// TakeReading is used for a request to a sensor to take a reading.
	TakeReading
	// TemperatureValue is used to signify a temperature reading.
	TemperatureValue
	// SwitchOn signifies a device should switch on.
	SwitchOn
	// SwitchOff signifies a device should switch off.
	SwitchOff
	// Error signifies an Error Message.
	Error
)

// Message is sent between different devices.
type Message struct {
	//MessageType is the type of message.
	MessageType MessageType
	//Message is the message to process.
	Message interface{}
}
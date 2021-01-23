// Package actors is used to hold common actors that can be utilized from outside
// packages.
package actors

import (
	"time"

	"github.com/MichaelS11/go-dht"
	"github.com/jmross14/go-smarthome-common/v2/structs"
)

// DHT22 is a model of what is needed to use a DHT22 Temperature Sensor.
type DHT22 struct {
	// actionChan is the channel messages are sent to the actor loop on.
	actionChan chan action
	// Reading contains a reading from the sensor
	reading structs.TemperatureReading
	// dht reads from the sensor
	dht *dht.DHT
}

// StartActor starts the actor and initializes its variables.
func (dht22 *DHT22) StartActor() *DHT22 {
	ch := make(chan action, 16)

	dht.HostInit()
	dht, _ := dht.NewDHT("GPIO4", dht.Celsius, "")

	humidity, temperature, _ := dht.ReadRetry(15)
	sensor := DHT22{
		ch,
		structs.TemperatureReading{Humidity: humidity, Temperature: temperature},
		dht,
	}

	go sensor.actorLoop(sensor.actionChan)
	return &sensor
}

// Tell sends a message to the actor and requests an operation that requires no response.
func (dht22 *DHT22) Tell(msg structs.Message) {
	//Intentionally left empty
}

// Ask sends a message to the actor and requests an operation that requires a response.
func (dht22 *DHT22) Ask(msg structs.Message) structs.Message {
	if msg.MessageType == structs.TakeReading {
		ch := make(chan structs.TemperatureReading, 1)
		dht22.actionChan <- func() {
			ch <- dht22.reading
		}
		reading := <-ch

		return structs.Message{MessageType: structs.TakeReading, Message: reading}
	} else {
		return structs.Message{MessageType: structs.Error, Message: "Unsupported Message"}
	}
}

// takeTemperature sets reading in the temperature sensor struct from the physical
// sensor
func (dht22 *DHT22) takeTemperatureReading() {

	humidity, temperature, err := dht22.dht.Read()
	if err == nil {
		dht22.reading = structs.TemperatureReading{Humidity: humidity, Temperature: temperature}
	}
}

// actorLoop runs indefinitely and processes requests.
func (dht22 *DHT22) actorLoop(ch <-chan action) {
	//Used to take the temperature every 10 seconds due to dht.Read() failing sometimes.
	//Do not want to block for the 10-12 seconds for dht.ReadRetry().
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case action := <-ch:
			action()
		case <-ticker.C:
			dht22.takeTemperatureReading()
		}
	}
}

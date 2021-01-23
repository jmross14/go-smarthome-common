// Package actors is used to hold common actors that can be utilized from outside
// packages.
package actors

import (
	"github.com/jmross14/go-smarthome-common/v2/structs"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

// Relay is a model of what is needed to use a relay.
type Relay struct {
	actionChan 	chan action
	pin			gpio.PinIO
}

// StartActor starts the actor and initializes its variables.
func (relay *Relay) StartActor() *Relay {
	ch := make(chan action, 16)

	host.Init()
	relay = &Relay{ch, gpioreg.ByName("GPIO17")}
	relay.pin.Out(gpio.Low)
	go relay.actorLoop(ch)
	
	return relay
}

// Tell sends a message to the actor and requests an operation that requires no response.
func (relay *Relay) Tell(msg structs.Message) {
	if msg.MessageType == structs.SwitchOn {
		relay.relayOn()
	} else if msg.MessageType == structs.SwitchOff {
		relay.relayOff()
	}
}

// Ask sends a message to the actor and requests an operation that requires a response.
func (relay *Relay) Ask(msg structs.Message) structs.Message {
	return structs.Message{MessageType: structs.Error, Message: "Unsupported Operation"}
} 

// relayOn switches the relay on.
func(relay *Relay) relayOn() {
	relay.actionChan <- func() {
		relay.pin.Out(gpio.High)
	}
}

// relayOff switches the relay off.
func(relay *Relay) relayOff() {
	relay.actionChan <- func() {
		relay.pin.Out(gpio.Low)
	}
}

// actorLoop runs indefinitely and processes requests.
func(relay *Relay) actorLoop(ch <-chan action) {
	for {
		select {
		case action := <- ch:
			action()
		}
	}
}
// Package actors is used to hold common actors that can be utilized from outside
// packages.
package actors

import "github.com/jmross14/go-smarthome-common/v2/structs"

// Actor defines a common interface for all actors to follow.
type Actor interface {
	// StartActor starts the actor and initializes its variables.
	StartActor() interface {}
	// Tell sends a message to the actor and requests an operation that requires no response.
	Tell(structs.Message)
	// Ask sends a message to the actor and requests an operation that requires a response.
	Ask(structs.Message) structs.Message
	// actorLoop runs indefinitely and processes requests.
	actorLoop(ch <-chan action)
}

type action func()
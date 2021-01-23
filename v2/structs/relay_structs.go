package structs

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
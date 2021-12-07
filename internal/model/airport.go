package model

//Airport struct
type Airport struct {
	ID       uint64
	Name     string
	Location string
}

//EventType type of CRUD event (Created, Updated or Removed)
type EventType uint8

//EventStatus status of the event (Deferred or Processed)
type EventStatus uint8

const (
	_ EventType = iota
	//Created when Airport is created
	Created
	//Updated when some fields of Airport is updated
	Updated
	//Removed when Airport instances is marked as removed
	Removed
)
const (
	_ EventStatus = iota
	//Deferred when event is not yet processed
	Deferred
	//Processed when event is processed
	Processed
)

// AirportEvent - CRUD event of Airport records
type AirportEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Airport
}

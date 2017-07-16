package entities

// Dock describes a dock entity
type Dock struct {
	ID   string
	name string

	longitude string
	latitude  string
}

// Docks list has all the fixture data for docks. @todo these will be taken from the database
var Docks = []*Dock{
	{ID: "1", name: "Westerdoksdijk", latitude: "Hi", longitude: "hi"},
	{ID: "2", name: "Amsterdam Central", latitude: "222", longitude: "333"},
}

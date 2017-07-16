package main

type Dock struct {
	ID   string
	name string

	longitude string
	latitude  string
}

var docks = []*Dock{
	{ID: "1", name: "Westerdoksdijk", latitude: "Hi", longitude: "hi"},
	{ID: "2", name: "Amsterdam Central", latitude: "222", longitude: "333"},
}

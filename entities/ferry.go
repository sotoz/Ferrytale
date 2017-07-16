package entities

type Ferry struct {
	ID   string
	name string
}

var Ferries = []*Ferry{
	{ID: "1", name: "Westerdoksdijk"},
	{ID: "2", name: "Amsterdam Central"},
}

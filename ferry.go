package main

type Ferry struct {
	ID   string
	name string
}

var ferries = []*Ferry{
	{ID: "1", name: "Westerdoksdijk"},
	{ID: "2", name: "Amsterdam Central"},
}

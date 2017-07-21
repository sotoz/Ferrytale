package entities

type Ferry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Ferries = []*Ferry{
	{ID: "1", Name: "Westerdoksdijk"},
	{ID: "2", Name: "Amsterdam Central"},
}

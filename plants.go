package main

type Plant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var plants = []Plant{
	Plant{
		ID:   "001",
		Name: "Fiddle Leaf Fig",
	},
	Plant{
		ID:   "002",
		Name: "Swiss Cheese Plant",
	},
	Plant{
		ID:   "003",
		Name: "Macho Fern",
	},
	Plant{
		ID:   "004",
		Name: "ZZ Plant",
	},
}

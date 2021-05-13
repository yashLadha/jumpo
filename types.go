package main

type Jumpo struct {
	// Prefix to enter on the prompt to move to the following location.
	// Acts as a prefix for jumplist but on the console.
	Prefix string `json:"pref"`
	// Location that need to be used for jump-list prefix. Will be a valid
	// system path.
	Location string `json:"location"`
}

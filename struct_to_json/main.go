package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Breed    string
	WeightKg int   `json:"weightKg,omitempty"`
	Size     *Size `json:"size,omitempty"`
	Sex      *int
}

type Size struct {
	width  int
	height int
}

func main() {
	dog := Dog{
		Breed: "gutou",
		Sex:   nil,
	}

	m, err := json.Marshal(dog)
	if err == nil {
		fmt.Println(string(m))
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Bogart", "Bergman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"McQueen", "Bisset"}},
}

func main() {
	// marshaling
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("Marshaling nie powiódł się: %s", err)
	}
	fmt.Printf("%s\n", data)

	// unmarshaling
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Unmarshaling nie powiódł się: %s", err)
	}
	fmt.Println(titles)
}

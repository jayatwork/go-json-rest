//helper program first iteration movies

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Life..the real one in prison", Year: 1999, Color: true,
		Actors: []string{"Eddie Murphy", "Martin Lawrence"}},
	{Title: "Good, the Bad, the Ugly", Year: 1967, Color: false,
	        Actors: []string{"Lee Van cleef", "Clint Eastwood", "thatguyTuko"}},
	{Title: "The Hateful Eight", Year: 2016, Color: true,
	        Actors: []string{"Kirk Douglas", "Samuel L. Jackson"}},

	//more if needed
}

func main() {
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}

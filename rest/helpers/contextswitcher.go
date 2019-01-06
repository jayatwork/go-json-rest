/*helper program that schedules context switching to another
core work grup responsiblity withinn single buiness day*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type DailyWorker struct {
	Workgroup string
	Hours int `json:"timeinhours"`
	Scope bool `json:"scope,omitempty"`
	Contributors []string
	Exercise bool `json:"exercise,omitempty"` //will implement later for prior to work mindfulness
}

var defaultday = []DailyWorker{
	{Workgroup: "OCP PAASTeam", Hours: 6, Scope: true,
		Contributors: []string{"JasonA", "CoworkerA", "CoworkerB"}},
	{Workgroup: "DevopsCOE", Hours: 1, Scope: true,
	        Contributors: []string{"JasonA", "CoworkerA", "CoworkerB"}},
	{Workgroup: "LegacySupport", Hours: 1, Scope: true,
	        Contributors: []string{"JasonA", "CoworkerA", "CoworkerB"}},

	//more if needed
}

func main() {
	data, err := json.MarshalIndent(defaultday, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}

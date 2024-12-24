package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// * The fields should be PascalCase in order to export the fields
type person struct {
	// First     string
	First     string `json:"_,omitempty"` // * keep the same name and add omitempty
	Last      string `json:"last,omitempty"`
	age       int    // ! not exported
	HairColor string `json:"-"` // * fields that using "-" will be omitted
}

func main() {
	p1 := person{"Amir", "Zare", 27, "brown"}
	p2 := person{"Faeze", "Arefi", 24, "brown"}
	p3 := person{First: "New User"}

	people := []person{p1, p2, p3}

	fmt.Printf("people: %+v\n", people)

	// * marshal
	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal:", bs)         // []byte
	fmt.Println("marshal:", string(bs)) // string

	// * unmarshal
	var up []person
	err = json.Unmarshal(bs, &up)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unmarshal %+v\n", up)
}

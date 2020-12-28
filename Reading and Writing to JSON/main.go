package main

import (
	"fmt"
	"log"

	"account/personpb"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const (
	outputFileName string = "output.bin"
)

func main() {
	runJSON()
	runEmun()
}

func runEmun() {
	person := &personpb.Person{
		Age:          20,
		FirstName:    "Ichigo",
		LastName:     "Kurosaki",
		Height:       1.90,
		PhoneNumbers: []string{"2121312312", "312312312"},
		EyeColour:    personpb.Person_EYE_BROWN,
	}

	fmt.Println("Person created,", person)
	person.EyeColour = personpb.Person_EYE_GREEN
	fmt.Println("Person updated,", person)
}

func runJSON() {
	person := createPerson()
	json, err := writeJSON(person)

	if err != nil {
		log.Fatalln("Can't transform struct into json")
		return
	}

	fmt.Println("Struct transformed into json,", json)
	emptyPerson := &personpb.Person{}
	if err2 := readJSON(json, emptyPerson); err2 != nil {
		log.Fatalln("Can't transform json into struct")
		return
	}

	fmt.Println("Data read from json", emptyPerson)
}

func createPerson() proto.Message {
	person := &personpb.Person{
		Age:          18,
		FirstName:    "Jaimito",
		LastName:     "Perez",
		Height:       1.78,
		PhoneNumbers: []string{"342", "432"},
		EyeColour:    personpb.Person_EYE_GREEN,
	}

	return person
}

func writeJSON(data proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{}

	out, err := marshaler.MarshalToString(data)

	if err != nil {
		return "", err
	}

	return out, nil
}

func readJSON(json string, message proto.Message) error {
	return jsonpb.UnmarshalString(json, message)
}

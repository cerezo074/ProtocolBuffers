package main

import (
	"fmt"
	"log"

	"account/userpb"

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
	person := &userpb.Person{
		Age:          20,
		FirstName:    "Ichigo",
		LastName:     "Kurosaki",
		Height:       1.90,
		PhoneNumbers: []string{"2121312312", "312312312"},
		EyeColour:    userpb.Person_EYE_BROWN,
	}

	fmt.Println("Person created,", person)
	person.EyeColour = userpb.Person_EYE_GREEN
	fmt.Println("Person updated,", person)
}

func runJSON() {
	user := createUser()
	json, err := writeJSON(user)

	if err != nil {
		log.Fatalln("Can't transform struct into json")
		return
	}

	fmt.Println("Struct transformed into json,", json)
	emptyUser := &userpb.User{}
	if err2 := readJSON(json, emptyUser); err2 != nil {
		log.Fatalln("Can't transform json into struct")
		return
	}

	fmt.Println("Data read from json", emptyUser)
}

func createUser() *userpb.User {
	user := userpb.User{
		Firstname: "Jaimito",
		Lastname:  "Perez",
		Emails: []string{
			"jaimito@f.co",
			"perez@w.com",
		},
	}

	return &user
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

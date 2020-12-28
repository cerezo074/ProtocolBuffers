package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"account/personpb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	outputFileName string = "output.bin"
)

/*
This is the implementation based on google example from:
https://developers.google.com/protocol-buffers/docs/gotutorial
*/

func main() {
	peopleBook := createBook()
	writeFile(peopleBook, outputFileName)
	emptyPeopleBook := &personpb.AddressBook{}
	readFile(outputFileName, emptyPeopleBook)
	fmt.Println("Data read from file", outputFileName, emptyPeopleBook)
}

func createBook() proto.Message {
	person := personpb.Person{
		Name:        "Ichigo",
		Id:          2020,
		Email:       "email@coco.com",
		LastUpdated: timestamppb.Now(),
		Phones: []*personpb.Person_PhoneNumber{
			{Number: "22321", Type: personpb.Person_MOBILE},
			{Number: "31231", Type: personpb.Person_HOME},
		},
	}

	person2 := personpb.Person{
		Name:        "Kenpachi",
		Id:          2020,
		Email:       "gotei@coco.com",
		LastUpdated: timestamppb.Now(),
		Phones: []*personpb.Person_PhoneNumber{
			{Number: "321", Type: personpb.Person_MOBILE},
			{Number: "432", Type: personpb.Person_HOME},
		},
	}

	book := personpb.AddressBook{
		People: []*personpb.Person{
			&person,
			&person2,
		},
	}

	return &book
}

func writeFile(data proto.Message, filename string) {
	rawData, err := proto.Marshal(data)

	if err != nil {
		log.Fatalln("Can't prepare struct to raw data(bytes), ", err)
		return
	}

	if err2 := ioutil.WriteFile(filename, rawData, os.ModePerm); err2 != nil {
		log.Fatalln("Can't write data to filename ", filename, " with error, ", err)
		return
	}

	fmt.Println("Data was saved at", filename, "name")
}

func readFile(filename string, message proto.Message) {
	rawData, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Cant read file ", filename, " with the following error ", err)
		return
	}

	if err := proto.Unmarshal(rawData, message); err != nil {
		log.Fatalln("Cant fit file data into struct, with error ", err)
		return
	}
}

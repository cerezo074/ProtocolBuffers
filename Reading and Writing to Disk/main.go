package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"account/userpb"

	"google.golang.org/protobuf/proto"
)

const (
	outputFileName string = "output.bin"
)

func main() {
	user := createUser()
	writeFile(user, outputFileName)
	emptyUser := &userpb.User{}
	readFile(outputFileName, emptyUser)
	fmt.Println("Data read from file", outputFileName, emptyUser)
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

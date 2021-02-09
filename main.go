package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"os"

	pb "github.com/mattwelke/protobuf-test/generated/go"
	"google.golang.org/protobuf/proto"
)

func main() {
	fileMode := 0644

	// From https://developers.google.com/protocol-buffers/docs/gotutorial
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to Protobuf-encode person:", err)
	}
	if err := ioutil.WriteFile("person.pbdata", out, os.FileMode(fileMode)); err != nil {
		log.Fatalln("Failed to write person to disk:", err)
	}

	outJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to JSON-encode person:", err)
	}
	if err := ioutil.WriteFile("person.json", outJSON, os.FileMode(fileMode)); err != nil {
		log.Fatalln("Failed to write person JSON to disk:", err)
	}
}

package main

import (
	"fmt"
	"log"

	pb "github.com/mattwelke/protobuf-test/generated/go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	myMessage := &pb.MyMessage{
		MyBoolValue: &wrapperspb.BoolValue{
			Value: false,
		},
	}

	if myMessage.MyBoolValue != nil {
		fmt.Printf("Data provided. It is %v.\n", myMessage.MyBoolValue.Value)
	} else {
		fmt.Printf("Data not provided.\n")
	}

	out, err := proto.Marshal(myMessage)
	if err != nil {
		log.Fatalf("Failed to Protobuf-encode %v: %s", myMessage, err)
	}

	fmt.Printf("Proto-encoded message: %v\n", out)

	outJSON, err := protojson.Marshal(myMessage)
	if err != nil {
		log.Fatalf("Failed to JSON-encode %v: %s", myMessage, err)
	}

	fmt.Printf("JSON-encoded message: %s\n", string(outJSON))
}

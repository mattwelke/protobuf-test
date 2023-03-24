package main

import (
	"fmt"
	"log"

	pb "github.com/mattwelke/protobuf-test/generated/go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	myMessage := &pb.SearchRequest{
		Query:         "test query",
		PageNumber:    1,
		ResultPerPage: 1,
		Corpus:        pb.Corpus_CORPUS_IMAGES,
	}

	out, err := proto.Marshal(myMessage)
	if err != nil {
		log.Fatalf("Failed to Protobuf-encode %v: %s", myMessage, err)
	}

	fmt.Printf("Proto-encoded message: %v\n", out)
	fmt.Printf("# bytes = %d\n", len(out))

	outJSON, err := protojson.Marshal(myMessage)
	if err != nil {
		log.Fatalf("Failed to JSON-encode %v: %s", myMessage, err)
	}

	fmt.Printf("JSON-encoded message: %s\n", string(outJSON))
	fmt.Printf("# bytes = %d\n", len(outJSON))
}

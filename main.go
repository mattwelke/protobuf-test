package main

import (
	pb "github.com/mattwelke/protobuf-test/generated/go"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	// fileMode := 0644

	// item1 := &pb.Cart{
	// 	StringField: "some string",
	// 	FloatField:  1,
	// 	StringValue: &wrapperspb.StringValue{Value: "some other string"},
	// 	FloatValue:  &wrapperspb.FloatValue{Value: 2},
	// 	Metadata: []*pb.MetadataItem{
	// 		{
	// 			Key:   "some key",
	// 			Value: "some value",
	// 		},
	// 	},
	// }

	// item2 := &pb.Cart{
	// 	StringField: "some string",
	// 	FloatField:  1,
	// 	StringValue: &wrapperspb.StringValue{Value: "some other string"},
	// 	FloatValue:  &wrapperspb.FloatValue{Value: 2},
	// 	Metadata:    []*pb.MetadataItem{},
	// }

	myMessage := &pb.MyMessage{
		MyBoolValue: &wrapperspb.BoolValue{
			Value: false,
		},
	}

	if myMessage.MyBoolValue != nil {
		// fmt.Printf("Data provided. It is %v.\n", myMessage.MyBoolValue.Value)
	} else {
		// fmt.Printf("Data not provided.\n")
	}

	// itemName := "cart"

	// out, err := proto.Marshal(item)
	// if err != nil {
	// 	log.Fatalf("Failed to Protobuf-encode %s: %v", itemName, err)
	// }
	// if err := ioutil.WriteFile(fmt.Sprintf("%s.pbdata", itemName), out, os.FileMode(fileMode)); err != nil {
	// 	log.Fatalf("Failed to write %s to disk: %v", itemName, err)
	// }

	// outJSON, err := protojson.Marshal(item)
	// if err != nil {
	// 	log.Fatalf("Failed to JSON-encode %s: %v", itemName, err)
	// }
	// if err := ioutil.WriteFile(fmt.Sprintf("%s.json", itemName), outJSON, os.FileMode(fileMode)); err != nil {
	// 	log.Fatalf("Failed to write %s JSON to disk: %v", itemName, err)
	// }

	// fmt.Printf("item1 has metadata = %t\n", len(item1.Metadata) > 0)

	// fmt.Printf("item2 has metadata = %t\n", len(item2.Metadata) > 0)
}

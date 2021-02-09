# protobuf-test

Using Protobuf with more than one programming language. Uses Go and Java.

Generated code is in `generated/<lang>`. Serialized data is in `serialized`.

Generate types using `./proto_gen.sh`. This script generates both the Go and Java types.

Run the program, which saves data to disk (Go only for now) with `go run main.go`. Data will be serialized to Protobuf binary data and JSON in the current directory.

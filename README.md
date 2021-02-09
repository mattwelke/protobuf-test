# protobuf-test

Using Protobuf with more than one programming language. Examples in Go and Java.

Generated code is in `generated/<lang>`. Serialized data is in `serialized`.

To generate types, run `./proto_gen.sh`. This script generates both the Go and Java types. The Go types are used by the `main.go` test program in this repo, but the Java types are ununsed. They can be copy and pasted into a Java project to use them. The runtime Protobuf dependency must be added to such a Java project, like:

```xml
<dependency>
  <groupId>com.google.protobuf</groupId>
  <artifactId>protobuf-java</artifactId>
  <version>4.0.0-rc-2</version>
</dependency>
```

To run the test Go program (`main.go`), which saves data to disk, run `go run main.go`. Data will be serialized to Protobuf binary data and JSON in the current directory.

## Default value vs. not provided optional value

Sometimes you have data that isn't required. A client can provide it or not. In that situation, you need the server to be able to tell whether or not the client chose to provide the data.

With Protobuf, scalar value types (strings, ints, bools) are serialized to their default value if they aren't set. This means that the server has no way to differentiate between a client choosing to set the value to its default (like `0` for ints) and a client choosing to not set the value.

To solve this problem, Google created a well-known types package "wrappers". It has message types defined that are meant to be used to wrap scalars, like `google.protobuf.BoolValue` for booleans.

Using this type looks like this:

```proto3
message MyMessage {
  google.protobuf.BoolValue my_bool_value = 1;
}
```

### Creating messages

Then, the generated code lets you wrap your scalar inside the language-specific data structure in the client code.

Go:

```go
package main

import (
	pb "<path_to_generated_go_code>"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	withDataProvided := &pb.MyMessage{
		MyBoolValue: &wrapperspb.BoolValue{
			Value: false,
		},
	}

	withoutDataProvided := &pb.MyMessage{}
}
```

Java:

```java
package mypackage;

import com.google.protobuf.BoolValue;

import <path_to_generated_java_code>.MyMessage;

public class App {
    public static void main( String[] args )
    {
        MyMessage withDataProvided = MyMessage.newBuilder()
                .setMyBoolValue(BoolValue.newBuilder().setValue(false)).build();

        MyMessage withoutDataProvided = MyMessage.newBuilder().build();
    }
}
```

### Reading messages

Also, the generated code gives you mechanisms the server can use to detect whether the client provided the data.

Go:

```go
if myMessage.MyBoolValue != nil {
	fmt.Printf("Data provided. It is %v.\n", myMessage.MyBoolValue.Value)
} else {
	fmt.Printf("Data not provided.\n")
}
```

Java:

```java
if (myMessage.hasMyBoolValue()) {
    System.out.println("Data was provided. It is " + myMessage.getMyBoolValue().getValue() + ".");
} else {
    System.out.println("Data not provided.");
}
```

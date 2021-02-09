#!/bin/bash

SRC_DIR="."

GO_DST_DIR="generated/go"
rm -r $GO_DST_DIR
mkdir -p $GO_DST_DIR
protoc -I=$SRC_DIR --go_out=$GO_DST_DIR $SRC_DIR/addressbook.proto
echo "Generated Go."

JAVA_DST_DIR="generated/java"
rm -r $JAVA_DST_DIR
mkdir -p $JAVA_DST_DIR
protoc -I=$SRC_DIR --java_out=$JAVA_DST_DIR $SRC_DIR/addressbook.proto
echo "Generated Java."

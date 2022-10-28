SRC_DIR="protobuf"
TARGET_DIR="lib"

rm -rf $TARGET_DIR/*

protoc -I=$SRC_DIR --go_out=$TARGET_DIR $SRC_DIR/*.proto

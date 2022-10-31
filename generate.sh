SRC_DIR="protobuf"
TARGET_DIR="."

rm -rf "lib"

protoc -I=$SRC_DIR --go_out=$TARGET_DIR $SRC_DIR/*.proto

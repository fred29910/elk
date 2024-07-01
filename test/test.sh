#!/bin/bash

script_dir=$(dirname "$0")
script_dir=$(cd "$script_dir" && pwd)
root_dir=$(cd $script_dir/../ && pwd)
# echo $script_dir
# echo $root_dir

# pre and install 
cd $root_dir/cmd/protoc-gen-msg-handle && go install ./
cd $script_dir


# DEBUG=true protoc --go_out=. --plugin=protoc-gen-msg-handle --msg-handle_out=. proto/test.proto
DEBUG=true protoc --go_out=.  --msg-handle_out=. proto/test.proto


go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest
protoc \
  --debug_out=".:." \
  proto/test.proto
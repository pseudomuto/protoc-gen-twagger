package main

//go:generate go build -o ./fixtures main.go
//go:generate protoc --plugin=protoc-gen-fix=./fixtures --proto_path=./protos --fix_out=. -I../ -I../options -I../vendor/github.com/googleapis/googleapis ./protos/doc.proto ./protos/todo.proto
//go:generate rm fixtures

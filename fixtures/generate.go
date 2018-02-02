package main

//go:generate go build -o ./fixtures main.go
//go:generate protoc --plugin=protoc-gen-fix=./fixtures -I=./protos --fix_out=. -I../ -I../options ./protos/doc.proto ./protos/todo/service.proto ./protos/greeter/service.proto
//go:generate rm fixtures

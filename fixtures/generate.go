package main

//go:generate protoc --descriptor_set_out=fileset.pb --include_imports --include_source_info -I. -I$GOPATH/src doc.proto todo/service.proto greeter/service.proto

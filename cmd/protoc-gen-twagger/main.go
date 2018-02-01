package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"

	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
)

func main() {
	req, err := parseCodeRequest(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	p := internal.NewPlugin(req)
	resp, err := p.Generate()
	if err != nil {
		log.Fatal(err)
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(data); err != nil {
		log.Fatal(err)
	}
}

func parseCodeRequest(r io.Reader) (*plugin_go.CodeGeneratorRequest, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		return nil, err
	}

	return req, nil
}

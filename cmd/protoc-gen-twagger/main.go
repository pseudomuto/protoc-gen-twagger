package main

import (
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"bytes"
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

	doc := internal.NewAPIDoc(req)
	buf := new(bytes.Buffer)

	if err = doc.ToJSON(buf); err != nil {
		log.Fatal(err)
	}

	resp := new(plugin.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
		Name:    proto.String("swagger.json"),
		Content: proto.String(buf.String()),
	})

	data, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(data); err != nil {
		log.Fatal(err)
	}
}

func parseCodeRequest(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		return nil, err
	}

	return req, nil
}

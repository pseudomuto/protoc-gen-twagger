package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"

	"io/ioutil"
	"os"
	"path"
)

func LoadCodeGenRequest() (*plugin_go.CodeGeneratorRequest, error) {
	return LoadCodeGenRequestWithGoPath(os.Getenv("GOPATH"))
}

func LoadCodeGenRequestWithGoPath(goPath string) (*plugin_go.CodeGeneratorRequest, error) {
	filePath := path.Join(goPath, "src", "github.com", "pseudomuto", "protoc-gen-twagger", "fixtures", "codegen.req")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		return nil, err
	}

	return req, nil
}

package internal

import (
	"github.com/golang/protobuf/protoc-gen-go/plugin"

	"encoding/json"
	"io"

	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type APIDoc struct {
	req *plugin_go.CodeGeneratorRequest
}

func NewAPIDoc(req *plugin_go.CodeGeneratorRequest) *APIDoc {
	return &APIDoc{req}
}

func (doc *APIDoc) ToJSON(w io.Writer) error {
	fd := utils.FindDocDescriptor(options.E_Api, doc.req.GetProtoFile())
	file := NewFile(fd)

	api, err := file.Generate()
	if err != nil {

		return err
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(api)
}

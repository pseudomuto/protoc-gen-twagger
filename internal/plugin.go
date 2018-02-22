package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"

	"bytes"
	"context"
	"encoding/json"
)

const outputFile = "swagger.json"

type Plugin struct {
	req *plugin_go.CodeGeneratorRequest
}

func (p *Plugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	ctx := context.Background()
	gen := NewGenerator(protokit.ParseCodeGenRequest(r))
	if err := gen.Generate(ctx); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")

	if err := enc.Encode(gen.api); err != nil {
		return nil, err
	}

	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(outputFile),
		Content: proto.String(buf.String()),
	})

	return resp, nil
}

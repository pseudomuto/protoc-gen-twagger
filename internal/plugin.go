package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"

	"bytes"
	"context"
	"encoding/json"
)

type Plugin struct {
	req *plugin_go.CodeGeneratorRequest
}

func NewPlugin(req *plugin_go.CodeGeneratorRequest) *Plugin {
	return &Plugin{req}
}

func (p *Plugin) Generate() (*plugin_go.CodeGeneratorResponse, error) {
	ctx := context.Background()
	gen := NewGenerator(p.parseFiles())
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
		Name:    proto.String("swagger.json"),
		Content: proto.String(buf.String()),
	})

	return resp, nil
}

func (p *Plugin) parseFiles() []*protokit.FileDescriptor {
	files := make([]*protokit.FileDescriptor, len(p.req.GetProtoFile()))

	for i, pf := range p.req.GetProtoFile() {
		files[i] = protokit.ParseFile(pf)
	}

	return files
}

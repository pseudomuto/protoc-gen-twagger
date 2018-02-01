package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"

	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pseudomuto/protoc-gen-twagger/internal/parser"
	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type Plugin struct {
	req *plugin_go.CodeGeneratorRequest
}

func NewPlugin(req *plugin_go.CodeGeneratorRequest) *Plugin {
	return &Plugin{req}
}

func (p *Plugin) Generate() (*plugin_go.CodeGeneratorResponse, error) {
	doc, err := p.generateSwagger()
	if err != nil {
		return nil, err
	}

	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String("swagger.json"),
		Content: proto.String(doc),
	})

	return resp, nil
}

func (p *Plugin) generateSwagger() (string, error) {
	files := p.parseFiles()
	doc := findTopLevelOption(files)
	if doc == nil {
		return "", fmt.Errorf("Couldn't find api options in any of the files")
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")

	if err := enc.Encode(doc); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (p *Plugin) parseFiles() []*parser.File {
	files := make([]*parser.File, len(p.req.GetProtoFile()))

	for i, pf := range p.req.GetProtoFile() {
		files[i] = parser.ParseFile(pf)
	}

	return files
}

func findTopLevelOption(files []*parser.File) *options.OpenAPI {
	for _, file := range files {
		if ext, err := proto.GetExtension(file.GetOptions(), options.E_Api); err == nil {
			if api, ok := ext.(*options.OpenAPI); ok {
				info := api.GetInfo()
				if info.GetDescription() == "" {
					info.Description = file.GetDescription()
				}

				return api
			}
		}
	}

	return nil
}

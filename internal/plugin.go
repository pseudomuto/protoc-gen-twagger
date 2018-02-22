package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"

	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pseudomuto/protoc-gen-twagger/options"
)

const outputFile = "swagger.json"

type Plugin struct {
	req *plugin_go.CodeGeneratorRequest
}

func (p *Plugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	descriptors := protokit.ParseCodeGenRequest(r)
	api, err := findOpenAPIDoc(descriptors)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	for _, file := range descriptors {
		generateFile(ctx, api, file)
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")

	if err := enc.Encode(api); err != nil {
		return nil, err
	}

	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(outputFile),
		Content: proto.String(buf.String()),
	})

	return resp, nil
}

func generateFile(ctx context.Context, api *options.OpenAPI, f *protokit.FileDescriptor) {
	api.Tags = append(api.Tags, ServicesToTags(ctx, f.GetServices())...)

	for _, svc := range f.GetServices() {
		for _, method := range svc.GetMethods() {
			url := fmt.Sprintf("/twirp/%s/%s", svc.GetFullName(), method.GetName())
			api.Paths[url] = MethodToPath(ctx, method, svc.GetName())
		}
	}

	for _, m := range f.GetMessages() {
		api.Components.Schemas[m.GetName()] = MessageToSchema(ctx, m)
	}
}

func findOpenAPIDoc(files []*protokit.FileDescriptor) (*options.OpenAPI, error) {
	for _, file := range files {
		ext, err := proto.GetExtension(file.GetOptions(), options.E_Api)
		if err != nil {
			continue
		}

		api, ok := ext.(*options.OpenAPI)
		if !ok {
			return nil, fmt.Errorf("Couldn't convert to OpenAPI object")
		}

		api.Info.Description = file.GetComments().String()
		api.Components.Schemas = make(map[string]*options.Schema)
		api.Paths = make(map[string]*options.Path)

		return api, nil
	}

	return nil, fmt.Errorf("Couldn't find api options in any of the files")
}

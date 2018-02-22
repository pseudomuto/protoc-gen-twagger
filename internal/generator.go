package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/pseudomuto/protokit"

	"context"
	"fmt"

	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type Generator struct {
	files []*protokit.FileDescriptor
	api   *options.OpenAPI
}

func NewGenerator(files []*protokit.FileDescriptor) *Generator {
	return &Generator{files: files}
}

func (g *Generator) Generate(ctx context.Context) error {
	var err error

	g.api, err = findOpenAPIDoc(g.files)
	if err != nil {
		return err
	}

	for _, file := range g.files {
		generateFile(ctx, g.api, file)
	}

	return nil
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

package internal

import (
	"github.com/golang/protobuf/proto"

	"context"
	"fmt"

	"github.com/pseudomuto/protoc-gen-twagger/internal/parser"
	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type Generator struct {
	files []*parser.File
	api   *options.OpenAPI
}

func NewGenerator(files []*parser.File) *Generator {
	return &Generator{files: files}
}

func (g *Generator) Generate(ctx context.Context) error {
	var err error

	g.api, err = findOpenAPIDoc(g.files)
	if err != nil {
		return err
	}

	for _, file := range g.files {
		ext, err := proto.GetExtension(file.GetOptions(), options.E_Namespace)
		if err != nil {
			continue
		}

		ns, ok := ext.(*string)
		if !ok {
			return fmt.Errorf("Unable to parse the twagger namespace for: %v", file.GetName())
		}

		generateFile(WithNamespace(ctx, *ns), g.api, file)
	}

	return nil
}

func generateFile(ctx context.Context, api *options.OpenAPI, f *parser.File) {
	api.Tags = append(api.Tags, ServicesToTags(ctx, f.GetServices())...)

	for _, svc := range f.GetServices() {
		for _, method := range svc.GetMethods() {
			api.Paths[method.GetUrl()] = MethodToPath(ctx, method, svc.GetName())
		}
	}

	for _, m := range f.GetMessages() {
		api.Components.Schemas[m.GetName()] = MessageToSchema(ctx, m)
	}
}

func findOpenAPIDoc(files []*parser.File) (*options.OpenAPI, error) {
	for _, file := range files {
		ext, err := proto.GetExtension(file.GetOptions(), options.E_Api)
		if err != nil {
			continue
		}

		api, ok := ext.(*options.OpenAPI)
		if !ok {
			return nil, fmt.Errorf("Couldn't convert to OpenAPI object")
		}

		api.Info.Description = file.GetDescription()
		api.Components = &options.Components{Schemas: make(map[string]*options.Schema)}
		api.Paths = make(map[string]*options.Path)
		return api, nil
	}

	return nil, fmt.Errorf("Couldn't find api options in any of the files")
}

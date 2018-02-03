package internal

import (
	"context"
	"strings"

	"github.com/pseudomuto/protoc-gen-twagger/internal/parser"
	"github.com/pseudomuto/protoc-gen-twagger/options"
)

func ServicesToTags(ctx context.Context, svcs []*parser.Service) []*options.Tag {
	tags := make([]*options.Tag, len(svcs))

	for i, svc := range svcs {
		tags[i] = &options.Tag{
			Name:        svc.GetName(),
			Description: summarize(svc.GetDescription()),
		}
	}

	return tags
}

func MethodToPath(ctx context.Context, method *parser.ServiceMethod, tag string) *options.Path {
	path := &options.Path{
		Post: &options.Operation{
			Summary:     summarize(method.GetDescription()),
			Description: describe(method.GetDescription()),
			Tags:        []string{tag},
		},
	}

	return path
}

func summarize(str string) string {
	return strings.SplitN(str, "\n", 2)[0]
}

func describe(str string) string {
	parts := strings.SplitN(str, "\n", 2)
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}

	return ""
}

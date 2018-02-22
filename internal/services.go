package internal

import (
	"github.com/pseudomuto/protokit"

	"context"
	"fmt"
	"strings"

	"github.com/pseudomuto/protoc-gen-twagger/options"
)

// ServicesToTags creates swagger Tag objects based on service names and descriptions
func ServicesToTags(ctx context.Context, svcs []*protokit.ServiceDescriptor) []*options.Tag {
	tags := make([]*options.Tag, len(svcs))

	for i, svc := range svcs {
		tags[i] = &options.Tag{
			Name:        svc.GetName(),
			Description: summarize(svc.GetComments().String()),
		}
	}

	return tags
}

// MethodToPath generates a Path object for a service method. Since we're being twirp specific, only the POST method is
// defined.
func MethodToPath(ctx context.Context, method *protokit.MethodDescriptor, tag string) *options.Path {
	return &options.Path{
		Post: &options.Operation{
			Summary:     summarize(method.GetComments().String()),
			Description: describe(method.GetComments().String()),
			Tags:        []string{tag},
			RequestBody: &options.RequestBody{
				Description: "The body for this request",
				Content: map[string]*options.MediaType{
					"application/json": {
						Schema: &options.Schema{
							Ref: fmt.Sprintf("#/components/schemas/%s", shortName(method.GetInputType())),
						},
					},
				},
				Required: true,
			},
			Responses: map[string]*options.Response{
				"200": &options.Response{
					Description: "The successful response",
					Content: map[string]*options.MediaType{
						"application/json": {
							Schema: &options.Schema{
								Ref: fmt.Sprintf("#/components/schemas/%s", shortName(method.GetOutputType())),
							},
						},
					},
				},
			},
		},
	}
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

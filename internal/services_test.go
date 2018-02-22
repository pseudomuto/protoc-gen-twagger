package internal_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pseudomuto/protokit"
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
)

type ServicesTest struct {
	suite.Suite
}

func TestServices(t *testing.T) {
	suite.Run(t, new(ServicesTest))
}

func (assert *ServicesTest) TestServicesToTags() {
	service := &protokit.ServiceDescriptor{
		ServiceDescriptorProto: &descriptor.ServiceDescriptorProto{
			Name: proto.String("MyService"),
		},
		Comments: &protokit.Comment{
			Leading: "Summary here\n\nDescription here",
		},
	}

	tags := internal.ServicesToTags(context.Background(), []*protokit.ServiceDescriptor{service})
	assert.Len(tags, 1)
	assert.Equal("MyService", tags[0].GetName())
	assert.Equal("Summary here", tags[0].GetDescription())
}

func (assert *ServicesTest) TestMethodToPath() {
	method := &protokit.MethodDescriptor{
		MethodDescriptorProto: &descriptor.MethodDescriptorProto{
			Name:       proto.String("DoSomething"),
			InputType:  proto.String("Some"),
			OutputType: proto.String("Thing"),
		},
		Comments: &protokit.Comment{
			Leading: "Summary here\n\nDescription down here",
		},
	}

	path := internal.MethodToPath(context.Background(), method, "MyService")
	assert.Equal("Summary here", path.GetPost().GetSummary())
	assert.Equal("Description down here", path.GetPost().GetDescription())
	assert.Equal([]string{"MyService"}, path.GetPost().GetTags())
	assert.Len(path.GetPost().GetResponses(), 1)

	req := path.GetPost().GetRequestBody()
	assert.Equal("#/components/schemas/Some", req.GetContent()["application/json"].GetSchema().GetRef())

	resp := path.GetPost().GetResponses()["200"]
	assert.Equal("#/components/schemas/Thing", resp.GetContent()["application/json"].GetSchema().GetRef())
}

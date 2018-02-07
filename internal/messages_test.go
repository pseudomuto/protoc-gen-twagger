package internal_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
	"github.com/pseudomuto/protoc-parser"
)

type MessagesTest struct {
	suite.Suite
}

func TestMessages(t *testing.T) {
	suite.Run(t, new(MessagesTest))
}

func (assert *MessagesTest) TestMessageToSchema() {
	message := &parser.Message{
		DescriptorProto: &descriptor.DescriptorProto{
			Name: proto.String("MyMessage"),
		},
		Description: "My message description",
		Fields: []*parser.MessageField{
			{
				FieldDescriptorProto: &descriptor.FieldDescriptorProto{
					Name:     proto.String("IntField"),
					JsonName: proto.String("intField"),
					Type:     descriptor.FieldDescriptorProto_TYPE_INT32.Enum(),
				},
				Description: "REQUIRED: Integer field",
			},
			{
				FieldDescriptorProto: &descriptor.FieldDescriptorProto{
					Name:     proto.String("RefField"),
					JsonName: proto.String("refField"),
					Type:     descriptor.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
					TypeName: proto.String("OtherSchema"),
				},
				Description: "Reference field",
			},
			{
				FieldDescriptorProto: &descriptor.FieldDescriptorProto{
					Name:     proto.String("TimestampField"),
					JsonName: proto.String("timestampField"),
					Type:     descriptor.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
					TypeName: proto.String(".google.protobuf.Timestamp"),
				},
				Description: "Timestamp field",
			},
		},
	}

	ctx := internal.WithNamespace(context.Background(), "somens")
	schema := internal.MessageToSchema(ctx, message)

	assert.Equal("My message description", schema.GetDescription())
	assert.Len(schema.Properties, 3)
	assert.Equal([]string{"intField"}, schema.Required)

	prop := schema.Properties["intField"]
	assert.Equal("Integer field", prop.Description)
	assert.Equal("integer", prop.Type)
	assert.Equal("int32", prop.Format)
	assert.Len(prop.Properties, 0)
	assert.Empty(prop.Ref)

	prop = schema.Properties["refField"]
	assert.Equal("Reference field", prop.Description)
	assert.Equal("object", prop.Type)
	assert.Len(prop.Properties, 0)
	assert.Empty(prop.Format)
	assert.Equal(prop.Ref, "#/components/schemas/OtherSchema")

	prop = schema.Properties["timestampField"]
	assert.Equal("Timestamp field", prop.Description)
	assert.Equal("string", prop.Type)
	assert.Equal("date-time", prop.Format)
	assert.Len(prop.Properties, 0)
	assert.Empty(prop.Ref)
}

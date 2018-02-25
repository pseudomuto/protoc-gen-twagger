package internal

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pseudomuto/protokit"

	"context"
	"fmt"
	"strings"

	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type typeFormat struct {
	name   string
	format string
}

var typeMap = map[descriptor.FieldDescriptorProto_Type]typeFormat{
	descriptor.FieldDescriptorProto_TYPE_INT32:    {"integer", "int32"},
	descriptor.FieldDescriptorProto_TYPE_UINT32:   {"integer", "int32"},
	descriptor.FieldDescriptorProto_TYPE_SINT32:   {"integer", "int32"},
	descriptor.FieldDescriptorProto_TYPE_FIXED32:  {"integer", "int32"},
	descriptor.FieldDescriptorProto_TYPE_SFIXED32: {"integer", "int32"},
	descriptor.FieldDescriptorProto_TYPE_INT64:    {"integer", "int64"},
	descriptor.FieldDescriptorProto_TYPE_UINT64:   {"integer", "int64"},
	descriptor.FieldDescriptorProto_TYPE_SINT64:   {"integer", "int64"},
	descriptor.FieldDescriptorProto_TYPE_FIXED64:  {"integer", "int64"},
	descriptor.FieldDescriptorProto_TYPE_SFIXED64: {"integer", "int64"},

	descriptor.FieldDescriptorProto_TYPE_DOUBLE: {"number", "double"},
	descriptor.FieldDescriptorProto_TYPE_FLOAT:  {"number", "float"},

	descriptor.FieldDescriptorProto_TYPE_BOOL: {"boolean", ""},

	descriptor.FieldDescriptorProto_TYPE_BYTES:  {"string", "bytes"},
	descriptor.FieldDescriptorProto_TYPE_STRING: {"string", ""},
}

type schema interface {
	GetName() string
	GetType() descriptor.FieldDescriptorProto_Type
	GetTypeName() string
	GetDescription() string
	GetProperties() map[string]schema
}

type msgSchema struct {
	*protokit.Descriptor
}

func (m *msgSchema) GetDescription() string {
	return m.GetComments().String()
}

func (m *msgSchema) GetType() descriptor.FieldDescriptorProto_Type {
	return descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func (m *msgSchema) GetTypeName() string { return "object" }

func (m *msgSchema) GetProperties() map[string]schema {
	props := make(map[string]schema)

	for _, f := range m.GetMessageFields() {
		props[f.GetJsonName()] = &fieldSchema{f}
	}

	return props
}

type fieldSchema struct {
	*protokit.FieldDescriptor
}

func (f *fieldSchema) GetProperties() map[string]schema {
	return make(map[string]schema)
}

func (f *fieldSchema) GetDescription() string {
	return f.GetComments().String()
}

// MessageToSchema converts a descriptor into a schema object for swagger
func MessageToSchema(ctx context.Context, m *protokit.Descriptor) *options.Schema {
	return makeSchema(ctx, &msgSchema{m})
}

func makeSchema(ctx context.Context, s schema) *options.Schema {
	format := typeName(s.GetType())
	schema := &options.Schema{
		Type:        format.name,
		Format:      format.format,
		Description: strings.TrimSpace(strings.TrimPrefix(s.GetDescription(), "REQUIRED:")),
		Properties:  make(map[string]*options.Schema),
	}

	for name, sch := range s.GetProperties() {
		if strings.HasPrefix(sch.GetDescription(), "REQUIRED:") {
			schema.Required = append(schema.Required, name)
		}

		optSchema := makeSchema(ctx, sch)
		if optSchema.Type == "object" {
			if sch.GetTypeName() == ".google.protobuf.Timestamp" {
				optSchema.Type = "string"
				optSchema.Format = "date-time"
			} else {
				optSchema.Ref = fmt.Sprintf("#/components/schemas/%s", shortName(sch.GetTypeName()))
			}
		}

		schema.Properties[name] = optSchema
	}

	return schema
}

func typeName(typ descriptor.FieldDescriptorProto_Type) typeFormat {
	if t, ok := typeMap[typ]; ok {
		return t
	}

	return typeFormat{"object", ""}
}

func shortName(str string) string {
	parts := strings.Split(str, ".")
	return parts[len(parts)-1]
}

package utils_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

// adapted from options/annotations.pb.go to avoid circular reference
type OpenAPI struct {
	Openapi string `protobuf:"bytes,1,opt,name=openapi" json:"openapi,omitempty"`
}

type DescriptorsTest struct {
	suite.Suite
	descs []*descriptor.FileDescriptorProto
}

func TestDescriptors(t *testing.T) {
	suite.Run(t, new(DescriptorsTest))
}

func (assert *DescriptorsTest) SetupSuite() {
	req, err := utils.LoadCodeGenRequest()
	assert.NoError(err)

	assert.descs = req.GetProtoFile()
}

func (assert *DescriptorsTest) TestFindDocDescriptor() {
	// copied from options/annotations.pb.go to avoid circular reference
	extDesc := &proto.ExtensionDesc{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*OpenAPI)(nil),
		Field:         81098,
		Name:          "com.pseudomuto.protoc_gen_twagger.options.api",
		Tag:           "bytes,81098,opt,name=api",
		Filename:      "annotations.proto",
	}

	assert.NotNil(utils.FindDocDescriptor(extDesc, assert.descs))
	assert.Nil(utils.FindDocDescriptor(extDesc, assert.descs[:1]))
}

func (assert *DescriptorsTest) TestFindFileDescriptor() {
	tests := []struct {
		name  string
		found bool
	}{
		{"todo.proto", true},
		{"doc.proto", true},
		{"whodis.proto", false},
	}

	for _, test := range tests {
		fd := utils.FindFileDescriptor(test.name, assert.descs)
		assert.Equal(test.found, fd != nil)
	}
}

func (assert *DescriptorsTest) TestFindServiceDescriptor() {
	fd := utils.FindFileDescriptor("todo.proto", assert.descs)
	assert.NotNil(fd)

	tests := []struct {
		name  string
		found bool
	}{
		{"Todo", true},
		{"whodis", false},
	}

	for _, test := range tests {
		sd := utils.FindServiceDescriptor(test.name, fd)
		assert.Equal(test.found, sd != nil)
	}
}

package utils_test

import (
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

package internal

import (
	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"fmt"

	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type File struct {
	*descriptor.FileDescriptorProto
	comments Comments
}

func NewFile(fd *descriptor.FileDescriptorProto) *File {
	return &File{
		FileDescriptorProto: fd,
		comments:            ParseComments(fd),
	}
}

func (f *File) Generate() (*options.OpenAPI, error) {
	if f.GetOptions() == nil {
		return nil, fmt.Errorf("%+v", f)
	}

	ext, err := proto.GetExtension(f.GetOptions(), options.E_Api)
	if err != nil {
		return nil, err
	}

	api, ok := ext.(*options.OpenAPI)
	if !ok {
		return nil, fmt.Errorf("Boom")
	}

	info := api.GetInfo()
	if info.GetDescription() == "" {
		info.Description = f.PackageDescription()
	}

	return api, nil
}

func (f *File) PackageDescription() string {
	return f.comments["2"]
}

package utils

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func FindFileDescriptor(name string, descs []*descriptor.FileDescriptorProto) *descriptor.FileDescriptorProto {
	for _, fd := range descs {
		if fd.GetName() == name {
			return fd
		}
	}

	return nil
}

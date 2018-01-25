package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func FindDocDescriptor(opt *proto.ExtensionDesc, descs []*descriptor.FileDescriptorProto) *descriptor.FileDescriptorProto {
	for _, fd := range descs {
		if _, err := proto.GetExtension(fd.GetOptions(), opt); err == nil {
			return fd
		}
	}

	return nil
}

func FindFileDescriptor(name string, descs []*descriptor.FileDescriptorProto) *descriptor.FileDescriptorProto {
	for _, fd := range descs {
		if fd.GetName() == name {
			return fd
		}
	}

	return nil
}

func FindServiceDescriptor(name string, fd *descriptor.FileDescriptorProto) *descriptor.ServiceDescriptorProto {
	for _, svc := range fd.GetService() {
		if svc.GetName() == name {
			return svc
		}
	}

	return nil
}

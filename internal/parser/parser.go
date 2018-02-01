package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"fmt"
)

const (
	// tag numbers in FileDescriptorProto
	packageCommentPath = 2
	messageCommentPath = 4
	serviceCommentPath = 6

	// tag numbers in DescriptorProto
	messageFieldCommentPath = 2

	// tag numbers in ServiceDescriptorProto
	serviceMethodCommentPath = 2
)

func ParseFile(fd *descriptor.FileDescriptorProto) *File {
	comments := ParseComments(fd)

	file := &File{FileDescriptorProto: fd, description: comments[fmt.Sprintf("%d", packageCommentPath)]}
	file.messages = parseMessages(fd.GetMessageType(), comments)
	file.services = parseServices(fd.GetService(), comments)

	return file
}

func parseMessages(protos []*descriptor.DescriptorProto, comments Comments) []*Message {
	msgs := make([]*Message, len(protos))

	for i, md := range protos {
		commentPath := fmt.Sprintf("%d.%d", messageCommentPath, i)

		msgs[i] = &Message{
			DescriptorProto: md,
			description:     comments[commentPath],
			fields:          parseMessageFields(md.GetField(), comments, commentPath),
		}
	}

	return msgs
}

func parseMessageFields(protos []*descriptor.FieldDescriptorProto, comments Comments, commentPrefix string) []*MessageField {
	fields := make([]*MessageField, len(protos))

	for i, fd := range protos {
		fields[i] = &MessageField{
			FieldDescriptorProto: fd,
			description:          comments[fmt.Sprintf("%s.%d.%d", commentPrefix, messageFieldCommentPath, i)],
		}
	}

	return fields
}

func parseServices(protos []*descriptor.ServiceDescriptorProto, comments Comments) []*Service {
	svcs := make([]*Service, len(protos))

	for i, sd := range protos {
		commentPath := fmt.Sprintf("%d.%d", serviceCommentPath, i)

		svcs[i] = &Service{
			ServiceDescriptorProto: sd,
			description:            comments[commentPath],
			methods:                parseServiceMethods(sd.GetMethod(), comments, commentPath),
		}
	}

	return svcs
}

func parseServiceMethods(protos []*descriptor.MethodDescriptorProto, comments Comments, commentPrefix string) []*ServiceMethod {
	methods := make([]*ServiceMethod, len(protos))

	for i, md := range protos {
		methods[i] = &ServiceMethod{
			MethodDescriptorProto: md,
			description:           comments[fmt.Sprintf("%s.%d.%d", commentPrefix, serviceMethodCommentPath, i)],
		}
	}

	return methods
}

package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"context"
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
	ctx := ContextWithComments(context.Background(), comments)
	ctx = ContextWithPackage(ctx, fd.GetPackage())

	file := &File{
		FileDescriptorProto: fd,
		Description:         comments[fmt.Sprintf("%d", packageCommentPath)],
	}

	file.Messages = parseMessages(ctx, fd.GetMessageType())
	file.Services = parseServices(ctx, fd.GetService())

	return file
}

func parseMessages(ctx context.Context, protos []*descriptor.DescriptorProto) []*Message {
	msgs := make([]*Message, len(protos))
	comments, _ := CommentsFromContext(ctx)
	pkg, _ := PackageFromContext(ctx)

	for i, md := range protos {
		commentPath := fmt.Sprintf("%d.%d", messageCommentPath, i)
		subCtx := ContextWithLocationPrefix(ctx, commentPath)

		msgs[i] = &Message{
			DescriptorProto: md,
			Description:     comments[commentPath],
			Fields:          parseMessageFields(subCtx, md.GetField()),
			Package:         pkg,
		}
	}

	return msgs
}

func parseMessageFields(ctx context.Context, protos []*descriptor.FieldDescriptorProto) []*MessageField {
	fields := make([]*MessageField, len(protos))
	comments, _ := CommentsFromContext(ctx)
	commentPrefix, _ := LocationPrefixFromContext(ctx)

	for i, fd := range protos {
		fields[i] = &MessageField{
			FieldDescriptorProto: fd,
			Description:          comments[fmt.Sprintf("%s.%d.%d", commentPrefix, messageFieldCommentPath, i)],
		}
	}

	return fields
}

func parseServices(ctx context.Context, protos []*descriptor.ServiceDescriptorProto) []*Service {
	svcs := make([]*Service, len(protos))
	comments, _ := CommentsFromContext(ctx)
	pkg, _ := PackageFromContext(ctx)

	for i, sd := range protos {
		commentPath := fmt.Sprintf("%d.%d", serviceCommentPath, i)
		subCtx := ContextWithLocationPrefix(ctx, commentPath)
		subCtx = ContextWithService(subCtx, sd.GetName())

		svcs[i] = &Service{
			ServiceDescriptorProto: sd,
			Description:            comments[commentPath],
			Methods:                parseServiceMethods(subCtx, sd.GetMethod()),
			Package:                pkg,
		}
	}

	return svcs
}

func parseServiceMethods(ctx context.Context, protos []*descriptor.MethodDescriptorProto) []*ServiceMethod {
	methods := make([]*ServiceMethod, len(protos))

	pkg, _ := PackageFromContext(ctx)
	svc, _ := ServiceFromContext(ctx)
	comments, _ := CommentsFromContext(ctx)
	commentPrefix, _ := LocationPrefixFromContext(ctx)

	for i, md := range protos {
		methods[i] = &ServiceMethod{
			MethodDescriptorProto: md,
			Description:           comments[fmt.Sprintf("%s.%d.%d", commentPrefix, serviceMethodCommentPath, i)],
			Url:                   fmt.Sprintf("/twirp/%s.%s/%s", pkg, svc, md.GetName()),
		}
	}

	return methods
}

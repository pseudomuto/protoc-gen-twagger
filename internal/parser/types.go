package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type File struct {
	*descriptor.FileDescriptorProto
	description string
	messages    []*Message
	services    []*Service
}

func (f *File) GetDescription() string  { return f.description }
func (f *File) GetMessages() []*Message { return f.messages }
func (f *File) GetServices() []*Service { return f.services }

func (f *File) GetMessage(name string) *Message {
	for _, m := range f.GetMessages() {
		if m.GetName() == name {
			return m
		}
	}

	return nil
}

func (f *File) GetService(name string) *Service {
	for _, s := range f.GetServices() {
		if s.GetName() == name {
			return s
		}
	}

	return nil
}

type Message struct {
	*descriptor.DescriptorProto
	description string
	fields      []*MessageField
}

func (m *Message) GetDescription() string            { return m.description }
func (m *Message) GetMessageFields() []*MessageField { return m.fields }

func (m *Message) GetMessageField(name string) *MessageField {
	for _, f := range m.GetMessageFields() {
		if f.GetName() == name {
			return f
		}
	}

	return nil
}

type MessageField struct {
	*descriptor.FieldDescriptorProto
	description string
}

func (mf *MessageField) GetDescription() string { return mf.description }

type Service struct {
	*descriptor.ServiceDescriptorProto
	description string
	methods     []*ServiceMethod
}

func (s *Service) GetDescription() string       { return s.description }
func (s *Service) GetMethods() []*ServiceMethod { return s.methods }

type ServiceMethod struct {
	*descriptor.MethodDescriptorProto
	description string
}

func (m *ServiceMethod) GetDescription() string { return m.description }

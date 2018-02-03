package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Properties map[string]string

type File struct {
	*descriptor.FileDescriptorProto
	Namespace   string // TODO: Get rid of this
	Description string
	Messages    []*Message
	Services    []*Service

	Props Properties
}

func (f *File) GetNamespace() string    { return f.Namespace }
func (f *File) GetDescription() string  { return f.Description }
func (f *File) GetMessages() []*Message { return f.Messages }
func (f *File) GetServices() []*Service { return f.Services }
func (f *File) GetProps() Properties    { return f.Props }

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
	Description string
	Fields      []*MessageField
	Package     string

	Props Properties
}

func (m *Message) GetDescription() string            { return m.Description }
func (m *Message) GetMessageFields() []*MessageField { return m.Fields }
func (m *Message) GetPackage() string                { return m.Package }
func (m *Message) GetProps() Properties              { return m.Props }

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
	Description string

	Props Properties
}

func (mf *MessageField) GetDescription() string { return mf.Description }
func (mf *MessageField) GetProps() Properties   { return mf.Props }

type Service struct {
	*descriptor.ServiceDescriptorProto
	Description string
	Methods     []*ServiceMethod
	Package     string

	Props Properties
}

func (s *Service) GetDescription() string       { return s.Description }
func (s *Service) GetMethods() []*ServiceMethod { return s.Methods }
func (s *Service) GetPackage() string           { return s.Package }
func (s *Service) GetProps() Properties         { return s.Props }

type ServiceMethod struct {
	*descriptor.MethodDescriptorProto
	Description string
	Url         string

	Props Properties
}

func (m *ServiceMethod) GetDescription() string { return m.Description }
func (m *ServiceMethod) GetUrl() string         { return m.Url }
func (m *ServiceMethod) GetProps() Properties   { return m.Props }

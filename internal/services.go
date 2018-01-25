package internal

import (
	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	options "google.golang.org/genproto/googleapis/api/annotations"
)

type Service struct {
	*descriptor.ServiceDescriptorProto
	methods []*ServiceMethod
}

func (s *Service) GetParsedMethods() []*ServiceMethod {
	return s.methods
}

type ServiceMethod struct {
	*descriptor.MethodDescriptorProto
	url string
	tag string
}

func (m *ServiceMethod) GetURL() string {
	return m.url
}

func (m *ServiceMethod) GetTag() string {
	return m.tag
}

func ParseService(svc *descriptor.ServiceDescriptorProto) *Service {
	service := &Service{ServiceDescriptorProto: svc, methods: []*ServiceMethod{}}

	for _, method := range service.GetMethod() {
		m := &ServiceMethod{MethodDescriptorProto: method, tag: service.GetName()}

		if ext, err := proto.GetExtension(m.GetOptions(), options.E_Http); err == nil {
			if opts, ok := ext.(*options.HttpRule); ok {
				m.url = opts.GetPost()
			}
		}

		service.methods = append(service.methods, m)
	}

	return service
}

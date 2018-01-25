package internal_test

import (
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

type ServicesTest struct {
	suite.Suite
	service *internal.Service
}

func TestServices(t *testing.T) {
	suite.Run(t, new(ServicesTest))
}

func (assert *ServicesTest) SetupSuite() {
	req, err := utils.LoadCodeGenRequest()
	assert.NoError(err)

	pf := utils.FindFileDescriptor("todo.proto", req.GetProtoFile())
	assert.NotNil(pf)

	assert.service = internal.ParseService(utils.FindServiceDescriptor("Todo", pf))
}

func (assert *ServicesTest) TestMethodParsing() {
	methods := assert.service.GetParsedMethods()
	assert.Len(methods, 2)

	m := methods[0]
	assert.Equal("/twirp/com.pseudomuto.todo.v1.Todo/CreateList", m.GetURL())
	assert.Equal("Todo", m.GetTag())

	m = methods[1]
	assert.Equal("/twirp/com.pseudomuto.todo.v1.Todo/AddItem", methods[1].GetURL())
	assert.Equal("Todo", m.GetTag())
}

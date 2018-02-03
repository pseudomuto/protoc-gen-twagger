package parser_test

import (
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal/parser"
	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

var req *plugin_go.CodeGeneratorRequest

type ParserTest struct {
	suite.Suite
}

func TestParser(t *testing.T) {
	req, _ = utils.LoadCodeGenRequest()
	suite.Run(t, new(ParserTest))
}

func (assert *ParserTest) TestParseFile() {
	proto := utils.FindFileDescriptor("doc.proto", req.GetProtoFile())
	assert.NotNil(proto)

	file := parser.ParseFile(proto)
	assert.Contains(file.GetDescription(), "# The official documentation for the Todo API.\n\n")
	assert.Len(file.GetMessages(), 0)
	assert.Len(file.GetServices(), 0)
}

func (assert *ParserTest) TestParseFileServices() {
	proto := utils.FindFileDescriptor("todo/service.proto", req.GetProtoFile())
	assert.NotNil(proto)

	file := parser.ParseFile(proto)
	assert.Len(file.GetServices(), 1)
	assert.Nil(file.GetService("swingandamiss"))

	svc := file.GetService("Todo")
	assert.Contains(svc.GetDescription(), "A service for managing \"todo\" items.\n\n")
	assert.Equal("com.pseudomuto.todo.v1", svc.GetPackage())
	assert.Len(svc.GetMethods(), 2)

	m := svc.GetMethods()[0]
	assert.Equal("Create a new todo list", m.GetDescription())

	m = svc.GetMethods()[1]
	assert.Equal("/twirp/com.pseudomuto.todo.v1.Todo/AddItem", m.GetUrl())
	assert.Equal("Add an item to your list\n\nAdds a new item to the specified list.", m.GetDescription())
}

func (assert *ParserTest) TestParseFileMessages() {
	proto := utils.FindFileDescriptor("todo/service.proto", req.GetProtoFile())
	assert.NotNil(proto)

	file := parser.ParseFile(proto)
	assert.Len(file.GetMessages(), 6)
	assert.Nil(file.GetMessage("swingandamiss"))

	m := file.GetMessage("AddItemRequest")
	assert.Equal("A request message for adding new items.", m.GetDescription())
	assert.Equal("com.pseudomuto.todo.v1", m.GetPackage())
	assert.Len(m.GetMessageFields(), 3)
	assert.Nil(m.GetMessageField("swingandamiss"))

	f := m.GetMessageField("completed")
	assert.Equal("Whether or not the item is completed.", f.GetDescription())
}

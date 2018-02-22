package internal_test

import (
	"github.com/pseudomuto/protokit/utils"
	"github.com/stretchr/testify/suite"

	"encoding/json"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type PluginTest struct {
	suite.Suite
}

func TestPlugin(t *testing.T) {
	suite.Run(t, new(PluginTest))
}

func (assert *PluginTest) TestGenerate() {
	set, err := utils.LoadDescriptorSet("..", "fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(set, "doc.proto", "todo/service.proto", "greeter/service.proto")
	plugin := new(internal.Plugin)

	resp, err := plugin.Generate(req)
	assert.NoError(err)
	assert.Len(resp.GetFile(), 1)

	file := resp.GetFile()[0]
	assert.Equal("swagger.json", file.GetName())
	assert.Contains(file.GetContent(), "{\n  \"openapi\": \"3.0\",")
}

func (assert *PluginTest) TestGenerateNoDocs() {
	set, err := utils.LoadDescriptorSet("..", "fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(set, "todo/service.proto", "greeter/service.proto")
	plugin := new(internal.Plugin)

	resp, err := plugin.Generate(req)
	assert.Nil(resp)
	assert.EqualError(err, "Couldn't find api options in any of the files")
}

func (assert *PluginTest) TestGeneratedContent() {
	set, err := utils.LoadDescriptorSet("..", "fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(set, "doc.proto", "todo/service.proto", "greeter/service.proto")
	assert.Len(req.FileToGenerate, 3)
	plugin := new(internal.Plugin)

	resp, err := plugin.Generate(req)
	assert.NoError(err)

	var api options.OpenAPI
	assert.NoError(json.Unmarshal([]byte(resp.GetFile()[0].GetContent()), &api))

	assert.Equal("3.0", api.Openapi)
	assert.Equal("Todo API", api.Info.Title)
	assert.Equal("Apache 2.0", api.Info.License.Name)

	assert.Len(api.Servers, 1)
	assert.Len(api.Tags, 2)
	assert.Len(api.Paths, 4)
}

package internal_test

import (
	"github.com/stretchr/testify/suite"

	"encoding/json"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
	"github.com/pseudomuto/protoc-gen-twagger/options"
)

type PluginTest struct {
	suite.Suite
}

func TestPlugin(t *testing.T) {
	suite.Run(t, new(PluginTest))
}

func (assert *PluginTest) TestGenerate() {
	req, _ := utils.LoadCodeGenRequest()
	resp, err := internal.NewPlugin(req).Generate()
	assert.NoError(err)
	assert.Len(resp.GetFile(), 1)

	file := resp.GetFile()[0]
	assert.Equal("swagger.json", file.GetName())
	assert.Contains(file.GetContent(), "{\n  \"openapi\": \"3.0\",")
}

func (assert *PluginTest) TestGeneratedContent() {
	req, _ := utils.LoadCodeGenRequest()
	resp, err := internal.NewPlugin(req).Generate()
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

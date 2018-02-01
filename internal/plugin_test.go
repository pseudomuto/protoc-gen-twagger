package internal_test

import (
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
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

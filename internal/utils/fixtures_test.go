package utils_test

import (
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

type FixturesTest struct {
	suite.Suite
}

func TestFixtures(t *testing.T) {
	suite.Run(t, new(FixturesTest))
}

func (assert *FixturesTest) TestLoadCodeGenRequest() {
	req, err := utils.LoadCodeGenRequest()
	assert.NoError(err)

	assert.Contains(req.GetFileToGenerate(), "doc.proto")
	assert.Contains(req.GetFileToGenerate(), "greeter/service.proto")
	assert.Contains(req.GetFileToGenerate(), "todo/service.proto")
}

func (assert *FixturesTest) TestLoadCodeGenRequestNotFound() {
	req, err := utils.LoadCodeGenRequestWithGoPath("./")
	assert.Nil(req)
	assert.Error(err)
	assert.Contains(err.Error(), "no such file or directory")
}

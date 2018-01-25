package internal_test

import (
	"github.com/stretchr/testify/suite"

	"strings"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

type FileTest struct {
	suite.Suite
	file *internal.File
}

func TestFile(t *testing.T) {
	suite.Run(t, new(FileTest))
}

func (assert *FileTest) SetupSuite() {
	req, err := utils.LoadCodeGenRequest()
	assert.NoError(err)

	pf := utils.FindFileDescriptor("doc.proto", req.GetProtoFile())
	assert.NotNil(pf)

	assert.file = internal.NewFile(pf)
}

func (assert *FileTest) TestPackageDescription() {
	str := assert.file.PackageDescription()
	assert.True(strings.HasPrefix(str, "# The official documentation for the Todo API.\n"))
}

func (assert *FileTest) TestGenerate() {
	api, err := assert.file.Generate()
	assert.NoError(err)

	assert.Equal("3.0", api.GetOpenapi())

	info := api.GetInfo()
	assert.Equal("Todo API", info.GetTitle())
	assert.Empty(info.GetTermsOfService())
	assert.Equal(assert.file.PackageDescription(), info.GetDescription())
	assert.Equal("0.1.0", info.GetVersion())

	contact := info.GetContact()
	assert.Equal("Todo Team", contact.GetName())
	assert.Equal("team@todo.com", contact.GetEmail())
	assert.Empty(contact.GetUrl())

	license := info.GetLicense()
	assert.Equal("Apache 2.0", license.GetName())
	assert.Equal("http://www.apache.org/licenses/LICENSE-2.0.html", license.GetUrl())

	servers := api.GetServers()
	assert.Len(servers, 1)
	assert.Equal("http://localhost:8000", servers[0].GetUrl())
	assert.Equal("The local development server.", servers[0].GetDescription())
}

package parser_test

import (
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal/parser"
	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

type CommentsTest struct {
	suite.Suite
	comments parser.Comments
}

func TestComments(t *testing.T) {
	suite.Run(t, new(CommentsTest))
}

func (assert *CommentsTest) SetupSuite() {
	req, err := utils.LoadCodeGenRequest()
	assert.NoError(err)

	pf := utils.FindFileDescriptor("todo/service.proto", req.GetProtoFile())
	assert.NotNil(pf)

	assert.comments = parser.ParseComments(pf)
}

func (assert *CommentsTest) TestComments() {
	tests := []struct {
		key   string
		value string
	}{
		{"6.0.2.1", "Add an item to your list\n\nAdds a new item to the specified list."}, // leading commend
		{"4.0.2.0", "The id of the list."},                                                // tailing comment
	}

	for _, test := range tests {
		assert.Equal(test.value, assert.comments[test.key])
	}
}

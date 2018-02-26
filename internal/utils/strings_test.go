package utils_test

import (
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal/utils"
)

type StringTest struct {
	suite.Suite
}

func TestString(t *testing.T) {
	suite.Run(t, new(StringTest))
}

func (assert *StringTest) TestFirstParagraph() {
	tests := map[string]string{
		"Some content\nOther content":       "Some content",
		"Some content\n\nOther content":     "Some content",
		"  Some content  \n\nOther content": "Some content",
	}

	for input, output := range tests {
		assert.Equal(output, utils.FirstParagraph(input))
	}
}

func (assert *StringTest) TestLastSubstring() {
	tests := map[string]string{
		"..com.package.v1.Message.Nested":   "Nested",
		"..com.package.v1.Message. Nested ": "Nested",
		" Nested":                           "Nested",
		"":                                  "",
		"  ":                                "",
	}

	for input, output := range tests {
		assert.Equal(output, utils.LastSubstring(input, "."))
	}
}

func (assert *StringTest) TestDescription() {
	tests := map[string]string{
		"REQUIRED: Some field comment": "Some field comment",
		" Comment ":                    "Comment",
		"":                             "",
		"  ":                           "",
	}

	for input, output := range tests {
		assert.Equal(output, utils.Description(input))
	}
}

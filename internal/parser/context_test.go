package parser_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal/parser"
)

type ContextTest struct {
	suite.Suite
}

func TestContext(t *testing.T) {
	suite.Run(t, new(ContextTest))
}

func (assert *ContextTest) TestContextWithComments() {
	ctx := context.Background()

	val, found := parser.CommentsFromContext(ctx)
	assert.Nil(val)
	assert.False(found)

	ctx = parser.ContextWithComments(ctx, make(parser.Comments))
	val, found = parser.CommentsFromContext(ctx)
	assert.NotNil(val)
	assert.True(found)
}

func (assert *ContextTest) TestContextWithLocationPrefix() {
	ctx := context.Background()

	val, found := parser.LocationPrefixFromContext(ctx)
	assert.Empty(val)
	assert.False(found)

	ctx = parser.ContextWithLocationPrefix(ctx, "prefix")
	val, found = parser.LocationPrefixFromContext(ctx)
	assert.Equal("prefix", val)
	assert.True(found)
}

func (assert *ContextTest) TestContextWithPackage() {
	ctx := context.Background()

	val, found := parser.PackageFromContext(ctx)
	assert.Empty(val)
	assert.False(found)

	ctx = parser.ContextWithPackage(ctx, "package")
	val, found = parser.PackageFromContext(ctx)
	assert.Equal("package", val)
	assert.True(found)
}

func (assert *ContextTest) TestContextWithService() {
	ctx := context.Background()

	val, found := parser.ServiceFromContext(ctx)
	assert.Empty(val)
	assert.False(found)

	ctx = parser.ContextWithService(ctx, "MyService")
	val, found = parser.ServiceFromContext(ctx)
	assert.Equal("MyService", val)
	assert.True(found)
}

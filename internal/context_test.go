package internal_test

import (
	"github.com/stretchr/testify/suite"

	"context"
	"testing"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
)

type ContextTest struct {
	suite.Suite
}

func TestContext(t *testing.T) {
	suite.Run(t, new(ContextTest))
}

func (assert *ContextTest) TestNamespace() {
	ctx := context.Background()

	ns, ok := internal.Namespace(ctx)
	assert.Empty(ns)
	assert.False(ok)

	ctx = internal.WithNamespace(ctx, "some-ns")
	ns, ok = internal.Namespace(ctx)
	assert.Equal("some-ns", ns)
	assert.True(ok)
}

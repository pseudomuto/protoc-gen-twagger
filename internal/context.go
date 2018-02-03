package internal

import (
	"context"
)

type contextKey string

const (
	namespaceContextKey = contextKey("namespace")
)

func WithNamespace(ctx context.Context, ns string) context.Context {
	return context.WithValue(ctx, namespaceContextKey, ns)
}

func Namespace(ctx context.Context) (string, bool) {
	ns, ok := ctx.Value(namespaceContextKey).(string)
	return ns, ok
}

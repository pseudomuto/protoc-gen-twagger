package parser

import (
	"context"
)

type contextKey string

const (
	commentsContextKey       = contextKey("comments")
	locationPrefixContextKey = contextKey("locationPrefix")
	packageContextKey        = contextKey("package")
	serviceContextKey        = contextKey("service")
)

func ContextWithComments(ctx context.Context, comments Comments) context.Context {
	return context.WithValue(ctx, commentsContextKey, comments)
}

func CommentsFromContext(ctx context.Context) (Comments, bool) {
	val, ok := ctx.Value(commentsContextKey).(Comments)
	return val, ok
}

func ContextWithLocationPrefix(ctx context.Context, locationPrefix string) context.Context {
	return context.WithValue(ctx, locationPrefixContextKey, locationPrefix)
}

func LocationPrefixFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(locationPrefixContextKey).(string)
	return val, ok
}

func ContextWithPackage(ctx context.Context, pkg string) context.Context {
	return context.WithValue(ctx, packageContextKey, pkg)
}

func PackageFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(packageContextKey).(string)
	return val, ok
}

func ContextWithService(ctx context.Context, pkg string) context.Context {
	return context.WithValue(ctx, serviceContextKey, pkg)
}

func ServiceFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(serviceContextKey).(string)
	return val, ok
}

package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"strconv"
	"strings"
)

type Comments map[string]string

func ParseComments(fd *descriptor.FileDescriptorProto) Comments {
	comments := make(Comments)

	for _, loc := range fd.GetSourceCodeInfo().GetLocation() {
		leading := loc.GetLeadingComments()
		trailing := loc.GetTrailingComments()

		if leading == "" && trailing == "" {
			continue
		}

		path := loc.GetPath()
		key := make([]string, len(path))
		for idx, p := range path {
			key[idx] = strconv.Itoa(int(p))
		}

		parts := make([]string, 0, 2)
		if leading != "" {
			parts = append(parts, scrub(leading))
		}

		if trailing != "" {
			parts = append(parts, scrub(trailing))
		}

		comments[strings.Join(key, ".")] = strings.Join(parts, "\x00")
	}

	return comments
}

func scrub(str string) string {
	return strings.TrimSpace(strings.Replace(str, "\n ", "\n", -1))
}

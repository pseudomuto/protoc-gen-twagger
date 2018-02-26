# protoc-gen-twagger

[![Travis Build Status][travis-svg]][travis-ci]
[![codecov][codecov-svg]][codecov-url]
[![Go Report Card][goreport-svg]][goreport-url]

A protobuf plugin that generates [Swagger] docs for your [Twirp] services.

Lots of assumptions being made at the moment...use with caution.

For a demo, clone the repo and run `make setup swagger`. Then open your browser to http://localhost:8080 and browse around.

## Running the Plugin

You'll need to install the binary.

`go get -u github.com/pseudomuto/protoc-gen-twagger/cmd/protoc-gen-twagger`

Even better, if you use a tool like [retool](https://github.com/twitchtv/retool) you can get a specific version. For
example:

`retool add github.com/pseudomuto/protoc-gen-twagger/cmd/protoc-gen-twagger <branch|tag>`

Once installed you can invoke it with `protoc` like so:

`protoc -I. -I$(GOPATH)src --twagger_out=. input_proto1.proto input_proto2.proto`

This will generate `swagger.json` in the current directory.

**Using Docker**

Alternatively you can run this with docker. There's a public image `pseudomuto/protoc-gen-twagger` which can be used to
run the plugin.

```shell
docker run --rm \
  -v $(pwd)/protos:/in:ro \
  -v $(pwd):/out:rw \
  pseudomuto/protoc-gen-twagger input_proto1.proto input_proto2.proto
```

Input paths are relative from `/in` inside the container. Take a look at the `run-docker` make task for a working
example.

[Swagger]: https://swagger.io/
[Twirp]: https://github.com/twitchtv/twirp
[travis-svg]:
  https://travis-ci.org/pseudomuto/protoc-gen-twagger.svg?branch=master
	"Travis CI build status SVG"
[travis-ci]:
  https://travis-ci.org/pseudomuto/protoc-gen-twagger
  "protoc-gen-twagger at Travis CI"
[codecov-svg]: https://codecov.io/gh/pseudomuto/protoc-gen-twagger/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/pseudomuto/protoc-gen-twagger
[goreport-svg]: https://goreportcard.com/badge/github.com/pseudomuto/protoc-gen-twagger
[goreport-url]: https://goreportcard.com/report/github.com/pseudomuto/protoc-gen-twagger

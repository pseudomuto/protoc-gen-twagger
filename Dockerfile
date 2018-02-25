# The builder image
FROM golang:1.10-alpine as builder
LABEL maintainer="pseudomuto <david.muto@gmail.com>"
WORKDIR /go/src/github.com/pseudomuto/protoc-gen-twagger
RUN apk --update add git make

COPY . ./
RUN make setup \
  && go build --ldflags="-s -w" ./cmd/protoc-gen-twagger

# The published image
FROM alpine:3.7
LABEL maintainer="pseudomuto <david.muto@gmail.com>"
RUN apk --update add protobuf-dev

COPY --from=builder /go/src/github.com/pseudomuto/protoc-gen-twagger/protoc-gen-twagger /usr/local/bin/
COPY --from=builder /go/src/github.com/pseudomuto/protoc-gen-twagger/options /go/src/github.com/pseudomuto/protoc-gen-twagger/options

VOLUME ["/in", "/out"]

WORKDIR /in
ENTRYPOINT ["protoc", "-I.", "-I/go/src", "--twagger_out=/out"]

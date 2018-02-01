.PHONY: setup test run swagger

TEST_DEPS = fixtures/codegen.req options/annotations.pb.go options/swagger.pb.go

setup:
	retool sync
	retool do dep ensure
	retool do govendor fetch github.com/googleapis/googleapis/google/api

clean:
	@rm -rf _output protoc-gen-twagger

build: cmd/protoc-gen-twagger/*.go internal/*.go options/*.go
	@echo Building protoc-gen-twagger
	@go build ./cmd/protoc-gen-twagger

run: build
	@echo Generating swagger.json...
	@mkdir -p _output
	@retool do protoc \
		--plugin=protoc-gen-twagger=./protoc-gen-twagger \
		-I. -I./options \
		-I./vendor/github.com/googleapis/googleapis \
		--proto_path=./fixtures/protos --twagger_out=./_output ./fixtures/protos/*.proto

fixtures/codegen.req: fixtures/protos/*.proto
	@echo Generating fixtures...
	@cd fixtures && go generate

options/annotations.pb.go options/swagger.pb.go: options/*.proto
	@echo Generating options...
	@cd options && go generate

test: $(TEST_DEPS)
	@echo Running tests...
	@go test -race -cover ./internal/...

swagger: $(TEST_DEPS) run
	@docker build -t twagger-test -f Dockerfile.test .
	@echo running Swagger UI at http://localhost:8080
	docker run --rm -it -p 8080:8080 twagger-test

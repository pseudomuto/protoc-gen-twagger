.PHONY: setup test test-ci run swagger

TEST_DEPS = fixtures/fileset.pb options/annotations.pb.go

setup:
	$(info Synching dev tools and dependencies...)
	@if test -z $(which retool); then go get github.com/twitchtv/retool; fi
	@retool sync
	@retool do dep ensure

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
		-I. -Ioptions -Ifixtures \
		--twagger_out=_output \
		fixtures/doc.proto fixtures/greeter/service.proto fixtures/todo/service.proto

fixtures/fileset.pb: fixtures/*.proto fixtures/**/*.proto
	@echo Generating fixtures...
	@cd fixtures && go generate

options/annotations.pb.go: options/*.proto
	@echo Generating options...
	@cd options && go generate

test: $(TEST_DEPS)
	@echo Running tests...
	@go test -race -cover ./internal/...

test-ci: $(TEST_DEPS)
	@retool do goverage -race -coverprofile=coverage.txt -covermode=atomic ./internal/...

swagger: $(TEST_DEPS) run
	@docker build -t twagger-test -f Dockerfile.test .
	@echo running Swagger UI at http://localhost:8080
	docker run --rm -it -p 8080:8080 twagger-test

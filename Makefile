.PHONY: setup test test-ci run swagger docker run-docker publish-docker

VERSION = "0.1.0-pre"

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
		--plugin=protoc-gen-twagger=protoc-gen-twagger \
		-I$(GOPATH)/src -Ifixtures \
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

docker: $(TEST_DEPS)
	@docker build -t pseudomuto/protoc-gen-twagger .

run-docker: docker
	$(info Running plugin with docker...)
	@mkdir -p _output/docker
	@docker run --rm \
		-v $(shell PWD)/fixtures:/in:ro \
		-v $(shell PWD)/_output/docker:/out:rw \
		pseudomuto/protoc-gen-twagger doc.proto greeter/service.proto todo/service.proto

swagger: $(TEST_DEPS) run
	@docker build -t twagger-test -f Dockerfile.test .
	@echo running Swagger UI at http://localhost:8080
	docker run --rm -it -p 8080:8080 twagger-test

release:
	@echo Releasing v${VERSION}...
	git add CHANGELOG.md Makefile
	git commit -m "Bump version to v${VERSION}"
	git tag -m "Version ${VERSION}" "v${VERSION}"
	git push && git push --tags

publish-docker:
	$(info Publishing docker image...)
	./publish "${VERSION}"

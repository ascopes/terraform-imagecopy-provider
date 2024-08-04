.DEFAULT_GOAL = help

.PHONY: build
build: tidy get vet
	go generate ./...
	go build ./...

.PHONY: clean
clean:
	go clean -i

.PHONY: get
get:
	go mod download -x

.PHONY: help
help:
	@-echo "Usage: make <target>"
	@-echo ""
	@-echo "Valid targets:"
	@-echo "    build - build the project."
	@-echo "    clean - delete any generated files in this project."
	@-echo "    get   - download any dependencies."
	@-echo "    help  - show this message."
	@-echo "    test  - run unit tests."
	@-echo "    tidy  - reformat code, update module descriptor."
	@-echo "    vet   - run any linters for the project."
	@-echo ""

.PHONY: test
test: build
	go test ./... -v

.PHONY: test
.PHONY: tidy
tidy: get
	go mod tidy
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

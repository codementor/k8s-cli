SHELL=/bin/bash -o pipefail
# Image URL to use all building/pushing image targets
CLI := kubectl-example
GIT_VERSION_PATH := github.com/codementor/k8s-cli/pkg/version.gitVersion
GIT_VERSION := $(shell git describe --abbrev=0 --tags | cut -b 2-)
GIT_COMMIT_PATH := github.com/codementor/k8s-cli/pkg/version.gitCommit
GIT_COMMIT := $(shell git rev-parse HEAD | cut -b -8)
SOURCE_DATE_EPOCH := $(shell git show -s --format=format:%ct HEAD)
BUILD_DATE_PATH := github.com/codementor/k8s-cli/pkg/version.buildDate
DATE_FMT := "%Y-%m-%dT%H:%M:%SZ"
BUILD_DATE := $(shell date -u -d "@$SOURCE_DATE_EPOCH" "+${DATE_FMT}" 2>/dev/null || date -u -r "${SOURCE_DATE_EPOCH}" "+${DATE_FMT}" 2>/dev/null || date -u "+${DATE_FMT}")
LDFLAGS := -X ${GIT_VERSION_PATH}=${GIT_VERSION} -X ${GIT_COMMIT_PATH}=${GIT_COMMIT} -X ${BUILD_DATE_PATH}=${BUILD_DATE}

export GO111MODULE=on

.PHONY: all
all: test

.PHONY: test
# Run tests
test:
	go test ./pkg/... ./cmd/... -v -mod=readonly -coverprofile cover.out

.PHONY: test-clean
# Clean test reports
test-clean:
	rm -f cover.out cover-integration.out

.PHONY: lint
lint:
ifeq (, $(shell which golangci-lint))
	./hack/install-golangcilint.sh
endif
	golangci-lint run

.PHONY: download
download:
	go mod download


.PHONY: cli-fast
# Build CLI but don't lint or run code generation first.
cli-fast:
	go build -ldflags "${LDFLAGS}" -o bin/${CLI} ./cmd/kubectl-example

.PHONY: cli
# Build CLI
cli: cli-fast

.PHONY: cli-clean
# Clean CLI build
cli-clean:
	rm -f bin/${CLI}

# Install CLI
cli-install:
	go install -ldflags "${LDFLAGS}" ./cmd/kubectl-example

.PHONY: clean
# Clean all
clean:  cli-clean test-clean


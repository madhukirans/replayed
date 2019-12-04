NAME:=replayed

DIST_DIR:=dist
BIN_DIR:=${DIST_DIR}/bin
BIN_NAME:=${NAME}
GO ?= go

.PHONY: all
all: build

#
# Go build related tasks
#
go-install:
	$(GO) install ./cmd/...

go-build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -o ${NAME}.linux.amd64  cmd/main.go

.PHONY: go-run
go-run: go-install
	$(GO) run ./cmd/main.go

.PHONY: go-fmt
go-fmt:
	gofmt -s -e -d -w $(shell find . -name "*.go" | grep -v /vendor/)

.PHONY: go-vet
go-vet:
	echo go vet $(shell go list ./... | grep -v /vendor/)

.PHONY: go-vendor
go-vendor:
	dep ensure -update

#
# Tests-related tasks
#
.PHONY: unit-test
unit-test: go-install
	go test -v ./cmd/...

.PHONY: code-cov
code-cov:
	go test -coverprofile -v ./...


.PHONY: prod-playbook
prod-playbook:
	ansible-playbook replayed.yaml -i inventories/prod/hosts

.PHONY: dev-playbook
dev-playbook:
	ansible-playbook replayed.yaml -i inventories/dev/hosts
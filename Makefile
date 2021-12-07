GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on
include .env

SERVICE_NAME=trv-airport-api
SERVICE_PATH=ozonmp/trv-airport-api

PGV_VERSION:="v0.6.1"
BUF_VERSION:="v0.56.0"

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

PYPI_USER=pypi_user
PYPI_PASSWORD=dsal2123!3d

ifeq ("NT", "$(findstring NT,$(OS_NAME))")
OS_NAME=Windows
endif

.PHONY: up
up:
	docker-compose up -d --no-deps trv-airport-api postgres mongo swagger-ui nginx jaeger prometheus grafana elasticsearch graylog
	cd migrations && goose postgres "host=localhost sslmode=disable dbname=$(POSTGRES_DB) port=5432 user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD)" up


.PHONY: build
build: build
	docker-compose down
	docker-compose build

.PHONY: up-build
up-build: build up
	echo "starting new containers"


.PHONY: call
call: migrate
	./curl_grpc.sh

.PHONY: hey
hey: migrate
	./hey.sh

.PHONY: run
run:
	go run cmd/grpc-server/main.go

.PHONY: lint
lint:
	golangci-lint run ./...


.PHONY: test
test:
	go test -v -race -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep total | awk '{print $$3}'


# ----------------------------------------------------------------

.PHONY: generate
generate: .generate-install-buf .generate-go .generate-python .generate-finalize-go .generate-finalize-python

.PHONY: generate
generate-go: .generate-install-buf .generate-go .generate-finalize-go

.generate-install-buf:
	@ command -v buf 2>&1 > /dev/null || (echo "Install buf" && \
    		mkdir -p "$(GO_BIN)" && \
    		curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(OS_NAME)-$(OS_ARCH)$(shell go env GOEXE) -o "$(BUF_EXE)" && \
    		chmod +x "$(BUF_EXE)")

.generate-go:
	$(BUF_EXE) generate

.generate-python:
	$(BUF_EXE) generate --template buf.gen.python.yaml

.generate-finalize-go:
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME)/* pkg/$(SERVICE_NAME)
	rm -rf pkg/$(SERVICE_NAME)/github.com/
	cd pkg/$(SERVICE_NAME) && ls go.mod || (go mod init github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME) && go mod tidy)

.generate-finalize-python:
	find pypkg/trv-airport-api -type d -exec touch {}/__init__.py \;

# ----------------------------------------------------------------

.PHONY: deps
deps: deps-go .deps-python

.PHONY: deps-go
deps-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest

.deps-python:
	python3 -m pip install grpcio-tools grpclib protobuf twine

.PHONY: build
build: generate .build

.PHONY: build-go
build-go: generate-go .build

.build:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/grpc-server$(shell go env GOEXE) ./cmd/grpc-server/main.go
#-----------------------------------------
.PHONY: pyupload
pyupload: .deps-python .generate-python
	docker-compose up -d pypi
	cd pypkg/trv-airport-api && python setup.py sdist
	twine upload  --username test --password hackme --repository-url http://localhost\:5000 pypkg/trv-airport-api/dist/*

.PHONY: pyinstall
pyinstall:
	pip install --extra-index-url http://localhost:5000 grpc-trv-airport-api

.PHONY: pycheck
pycheck: pyupload pyinstall
	python3 scripts/grpc_client.py

.PHONY: migrate
migrate:
	cd migrations && goose postgres "host=localhost sslmode=disable dbname=$(POSTGRES_DB) port=5432:5432 user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD)" up
.PHONY: migrate-down
migrate-down:
	cd migrations && goose postgres "host=localhost sslmode=disable dbname=$(POSTGRES_DB) port=5432:5432 user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD)" down

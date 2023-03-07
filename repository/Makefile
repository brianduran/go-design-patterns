BUILD_DATE = $(shell date +'%Y%m%d')
COMMIT_HASH = $(shell git rev-parse HEAD)
COMMIT_HASH_SHORT = $(shell echo $(COMMIT_HASH) | cut -c1-6)

LDFLAGS += -X "main.BuildDate=$(BUILD_DATE)"
LDFLAGS += -X "main.Commit=$(COMMIT_HASH_SHORT)"

PKGS ?= $(shell go list ./... | grep -v /vendor/ | grep -v /mocks)

CMD_NAME = "repository"

export GO111MODULE=on

.PHONY: all
all: build

.PHONY: build
build:
	@echo 'Building repository ($(COMMIT_HASH_SHORT))'
	@go build -o ./$(CMD_NAME) -ldflags '$(LDFLAGS)' ./cmd

.PHONY: build-linux
build-linux:
	@echo 'Building repository ($(COMMIT_HASH_SHORT))'
	CGO_ENABLED=0 GOOS=linux go build -o ./$(CMD_NAME) -ldflags '$(LDFLAGS)' ./cmd

.PHONY: build-image
build-image:
	@echo 'Building repository_pattern Docker image ($(COMMIT_HASH_SHORT))'
	@docker build -t $(CMD_NAME):$(COMMIT_HASH_SHORT) --build-arg COMMIT=$(COMMIT_HASH_SHORT) --build-arg GITHUB_USER=$(GITHUB_USER) --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) .

.PHONY: up
up:
	docker-compose up -d


.PHONY: check-all
check-all: check-style build test

.PHONY: check-style
check-style:
	@echo 'Checking code style'
	@ddgolint -set_exit_status ./...
	@golint -set_exit_status ./...

.PHONY: clean
clean:
	@rm tools.coverprofile || true
	@rm test_coverage_report.html || true
	@rm $(CMD_NAME) || true

.PHONY: init
init:
	git config core.hooksPath .githooks

.PHONY: test
test:
	@echo 'Running tests'
	@go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}go test \
		-test.timeout=120s -covermode=set \
		-coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile_part \
		{{.ImportPath}}{{end}}' $(PKGS) | \
		while read line; do bash -c "$$line" || exit 1; done
	@echo "mode: set" > tools.coverprofile
	@grep -h -v "^mode:" *.coverprofile_part >> "tools.coverprofile"
	@rm *.coverprofile_part
	@go tool cover -html=tools.coverprofile -o=test_coverage_report.html
	@echo "Overall coverage:"
	@go tool cover -func=tools.coverprofile | \
		tail -1 | \
		tr -s '	' | \
		cut -d '	' -f3
	@echo ''

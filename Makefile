PKGS ?= $(shell go list ./... | grep -v /vendor/ | grep -v /mocks)

export GO111MODULE=on

.PHONY: check-all
check-all: check-style test

.PHONY: check-style
check-style:
	@echo 'Checking code style'
	@ddgolint -set_exit_status ./...
	@golint -set_exit_status ./...

.PHONY: clean
clean:
	@rm tools.coverprofile || true
	@rm test_coverage_report.html || true

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

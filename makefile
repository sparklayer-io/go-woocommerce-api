.PHONY: fmt
# Format project code files
fmt:
	docker run -t --rm -v ./:/app -w /app golangci/golangci-lint:v1.64.6 golangci-lint run ./... --fix

.PHONY: lint
# Lint project code files
lint:
	docker run -t --rm -v ./:/app -w /app golangci/golangci-lint:v1.64.6 golangci-lint run ./...

.PHONY: tidy
# Sync dependencies and tidy the mod file
tidy:
	go mod tidy
	go mod vendor

.PHONY: help
# Show this help prompt
help:
	@echo
	@echo 'Usage:'
	@echo '	make <target> [flags...]'
	@echo ''
	@echo 'Targets:'
	@echo ''
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

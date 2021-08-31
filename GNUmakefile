default: build

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./internal/provider

build: fmt
	go install

generate: build
	go generate  ./...

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: fmt build generate testacc

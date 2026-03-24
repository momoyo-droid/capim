.PHONY: fmt vet lint vuln audit

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run --verbose

vuln:
	govulncheck ./...

audit: fmt vet lint vuln
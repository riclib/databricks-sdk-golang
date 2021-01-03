all : checks test

checks:
	go build all
	golangci-lint run

test: checks
	go test ./...

fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

deepcopy:
	./cmd/deepcopy-gen -i ./,./aws/...,./azure/... -h ./hack/boilerplate.go.txt -v 3

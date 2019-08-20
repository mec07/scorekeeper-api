.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -coverprofile=cover.out -covermode=atomic -coverpkg=./... ./...

.PHONY: cover
cover:
	go tool cover -html=cover.out

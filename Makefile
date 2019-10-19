.PHONY:test
test:
	go test -v -count=1 ./...

.PHONY: mocks
mocks:
	mockgen -source=./src/repositories/user.go -destination=./src/repositories/mocks/user.go
	mockgen -source=./src/providers/crypt.go -destination=./src/providers/mocks/crypt.go

.PHONY: run
run:
	go run main.go 
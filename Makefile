run:
	go run cmd/weatherapi/main.go

test:
	go test -v ./...

test-cov:
	go test -v ./... -coverprofile=c.out
	go tool cover -html=c.out
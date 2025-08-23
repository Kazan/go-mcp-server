run:
	go run cmd/server/main.go

tidy:
	@go mod tidy && go mod vendor

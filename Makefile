run:
	@go fmt ./...
	@swag init
	@go run ./
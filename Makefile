.PHONE: run
run:
	go run cmd/grpc-server/main.go

.PHONY: migration
migration:
	go run migrations/main.go
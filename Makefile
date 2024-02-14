all:
	swag fmt -g internal/ && swag init -g cmd/main.go && go run cmd/main.go

run:
	go run cmd/main.go

swagger:
	swag fmt -g internal/ && swag init -g cmd/main.go

fmt:
	swag fmt -g internal/
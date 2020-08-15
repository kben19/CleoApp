.PHONY: build

# go build command
build:
	@go build -v -o CleoApp ./main.go

# go run command
run:
	make build
	@./CleoApp
.PHONY: build

# go build command
build:
	@go build -v -o CleoApp ./main.go

# go run command
run:
	make build
	@./CleoApp

# dep ensure command
dep:
	dep ensure -v -vendor-only

# gcloud deploy command
deploy:
	@gcloud app deploy

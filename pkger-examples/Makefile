default: build
	go build -v -o example
linux:
	GOOS=linux go build -v -o example
	docker build -t pkger:example .
	docker run -p 3000:3000 pkger:example
build:
	find . -name ".DS_Store" -print -delete
	pkger
	go fmt ./...
	go mod tidy

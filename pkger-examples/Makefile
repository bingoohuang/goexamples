default:
	pkger
	go fmt ./...
	go mod tidy
	go build -v -o example
linux:
	pkger
	GOOS=linux go build -v -o example
	docker build -t pkger:example .
	docker run -p 3000:3000 pkger:example

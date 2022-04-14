up: docker-build docker-up

run: build go-run

test: build go-test

build: go-clean go-build

go-clean:
	go clean
	rm -rf main

go-build:
	go build ./src/main.go

go-run:
	go run ./src/main.go

go-test:
	go test -v ./...

docker-build:
	docker-compose --file ./docker/docker-compose.yml build --force

docker-up:
	docker-compose --file ./docker/docker-compose.yml up

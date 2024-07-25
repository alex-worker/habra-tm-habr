MAIN_PATH = "./cmd/app/main.go"

build: go-clean go-build

up: docker-build docker-up

run: build go-run

test: build go-test

clean: go-clean

go-clean:
	go clean
	rm -rf main

go-build:
	go build ${MAIN_PATH}

go-run:
	go run ${MAIN_PATH}

go-test:
	go test -v ./...

docker-build:
	docker-compose --file ./docker/docker-compose.yml build

docker-up:
	docker-compose --file ./docker/docker-compose.yml up

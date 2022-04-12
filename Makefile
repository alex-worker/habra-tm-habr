up: docker-build docker-up

run: go-run

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

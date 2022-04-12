up: docker-build docker-up

go-build:
	go build ./src/main.go

docker-build:
	docker-compose --file ./docker/docker-compose.yml build --force

docker-up:
	docker-compose --file ./docker/docker-compose.yml up

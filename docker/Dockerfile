FROM golang:1.18-alpine AS builder

WORKDIR /app

RUN apk add make

COPY go.mod ./
COPY go.sum ./

COPY Makefile ./

COPY ./src ./src

RUN make go-build

#ENTRYPOINT ["tail", "-f", "/dev/null"]

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/habra-tm-habr

CMD /app/habra-tm-habr

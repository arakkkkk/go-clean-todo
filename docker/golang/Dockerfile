FROM golang:1.21.0-alpine3.18

WORKDIR /
RUN apk update

RUN go install github.com/cosmtrek/air@latest

WORKDIR /go/src
# CMD ["go", "run", "cmd/app/main.go"]
CMD ["air", "-c", "air.toml"]


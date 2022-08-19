FROM golang:latest

WORKDIR /go/src/
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
COPY ./go.mod ./go.mod

RUN go mod tidy
RUN go build cmd/main.go
CMD ["./main"]
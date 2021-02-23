FROM golang:1.14.1-alpine as build-env

RUN mkdir -p /go/src/github.com/account-manager
WORKDIR /go/src/github.com/account-manager
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build *.go

CMD ["./main"]
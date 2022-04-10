FROM golang:1.16-alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go get -d -v ./....
RUN go install -v ./...

COPY *.go ./
ADD . .


RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]
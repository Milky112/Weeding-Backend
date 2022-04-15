FROM golang:1.16-alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

WORKDIR /app/go-app

COPY go.mod go.sum ./
RUN go mod download  && go mod verify

COPY . .


RUN go build -o ./out/go-app .

EXPOSE 8080
CMD ["./out/go-app"]
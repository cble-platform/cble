FROM golang:1.21-alpine

# Dependencies
RUN apk add git

RUN mkdir /app
ADD . /app/
WORKDIR /app

ENV PATH="${PATH}:/app"

RUN go mod download && go mod verify
RUN go build -o cble_server main.go

CMD ["./cble_server"]
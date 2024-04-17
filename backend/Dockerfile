FROM golang:1.22.2-alpine

# Dependencies
RUN apk add git

RUN mkdir /app
ADD . /app/
WORKDIR /app

ENV PATH="${PATH}:/app"

RUN go mod download && go mod verify
RUN go build -o cble_server .

CMD ["./cble_server"]
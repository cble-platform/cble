FROM golang:1.21-alpine

RUN mkdir /app
ADD . /app/
WORKDIR /app 

ENV PATH="${PATH}:/app"

RUN go mod download && go mod verify
RUN go build -o cble main.go

CMD ["./cble"]
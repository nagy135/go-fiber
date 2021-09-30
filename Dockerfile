FROM golang:1.17

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/gofiber/fiber/v2

CMD go run main.go

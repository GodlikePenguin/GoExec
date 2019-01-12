FROM golang:1.10-alpine

ADD main.go /app/main.go
WORKDIR /app
RUN go build main.go

CMD ["./main"]

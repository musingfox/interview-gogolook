FROM golang:1.18

WORKDIR /go/src/app
COPY ./src .

RUN go build -o main main.go

CMD ["./main"]

EXPOSE 5566

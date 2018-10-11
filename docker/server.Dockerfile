FROM golang:1.9.3-alpine3.7

COPY . $GOPATH/src/Frontend

WORKDIR $GOPATH/src/Frontend

RUN go build -o main main.go

ENTRYPOINT ["./main"]

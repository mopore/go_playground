FROM golang:1.21

WORKDIR /app

COPY go.mod /app
COPY main.go /app

RUN go get
RUN go build -o bin .

ENTRYPOINT ["/app/bin"]

FROM golang:1.22-alpine as builder
COPY . /app

WORKDIR /app

RUN go get -d -v ./...
RUN go build -o /app/main /app/main.go

FROM scratch
COPY --from=builder /app/main /app/main

CMD ["/app/main"]

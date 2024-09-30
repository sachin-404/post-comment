FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api cmd/api/main.go

FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /go/bin/api .

EXPOSE 8080

CMD ["./api"]
#build stage
FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o main .

#final stage
FROM alpine:3.21
ENV PORT=3001
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]


FROM golang:alpine AS builder
WORKDIR /app
COPY . . 
RUN go build -o services .

FROM alpine

WORKDIR /app
COPY --from=builder /app/services .
COPY .env .

CMD ["./services"]


FROM golang:1.17.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build .

FROM alpine
WORKDIR /app
COPY --from=builder /app/logger .
RUN ["./logger"]


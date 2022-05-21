FROM golang:1.17.5-bullseye AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build .

FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/logger .

EXPOSE ${PORT}

CMD ["./logger"]


FROM golang:1.18.3-bullseye AS builder
WORKDIR /app
COPY . .
RUN go build .

FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/worker .
EXPOSE ${PORT}

CMD ["./worker"]

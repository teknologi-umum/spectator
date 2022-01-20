FROM golang:1.17.5-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build .

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/worker .
ENV PORT=4444
EXPOSE 4444

CMD ["./worker"]

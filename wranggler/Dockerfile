FROM golang:1.17-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o /app/uwoh main.go

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/uwoh /app/uwoh
EXPOSE 4444
HEALTHCHECK CMD curl --fail http://localhost:4444/ping || exit 1
CMD ["/app/uwoh"]

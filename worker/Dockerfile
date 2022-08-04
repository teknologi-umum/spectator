FROM golang:1.18.5-bullseye AS builder

WORKDIR /app

COPY ./worker/ .

RUN go build .

FROM debian:bullseye

WORKDIR /app

RUN apt-get update && apt-get upgrade -y && apt-get install -y curl tar gzip

RUN curl -LO https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz && \
    tar -zxf grpcurl_1.8.6_linux_x86_64.tar.gz && \
    mv grpcurl /usr/bin/grpcurl && \
    chmod +x /usr/bin/grpcurl && \
    rm grpcurl_1.8.6_linux_x86_64.tar.gz LICENSE

COPY --from=builder /app/worker .

COPY ./proto/worker.proto .

EXPOSE ${PORT}

CMD ["./worker"]

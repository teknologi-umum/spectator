FROM golang:1.18.5-buster AS builder

WORKDIR /app

COPY ./video/ .

RUN go build -o video .

FROM debian:buster

WORKDIR /app

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl tar gzip build-essential pkg-config apt-utils debconf nasm && \
    apt-get install -y --install-suggests ffmpeg libvpx-dev libvpx6 vpx-tools libopus-dev \
    libx264-dev libbz2-dev liblzma-dev libvidstab-dev libx265-dev libnuma-dev libdav1d-dev

RUN curl -LO https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz && \
    tar -zxf grpcurl_1.8.6_linux_x86_64.tar.gz && \
    mv grpcurl /usr/bin/grpcurl && \
    chmod +x /usr/bin/grpcurl && \
    rm grpcurl_1.8.6_linux_x86_64.tar.gz LICENSE

COPY --from=builder /app/video /app/video

COPY ./proto/video.proto .

CMD [ "./video" ]

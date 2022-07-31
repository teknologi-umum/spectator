FROM golang:1.18.3-bullseye AS builder

WORKDIR /app

COPY . .

RUN go build -o video .

FROM debian:bullseye

WORKDIR /app

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl build-essential pkg-config && \
    apt-get install -y ffmpeg libvpx-dev libvpx6 vpx-tools libopus-dev libx264-dev libbz2-dev liblzma-dev libvidstab-dev libx265-dev

COPY --from=builder /app/video /app/video

CMD [ "./video" ]

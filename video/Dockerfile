FROM golang:1.18.3-bullseye AS builder

WORKDIR /app

COPY . .

RUN go build -o video .

FROM debian:bullseye

WORKDIR /app

RUN apt-get update && apt-get upgrade -y && apt-get install -y curl build-essential ffmpeg

COPY --from=builder /app/video /app/video

CMD [ "./video" ]

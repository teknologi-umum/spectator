FROM ruby:3.1.0-bullseye

WORKDIR /app

COPY . . 

RUN apt-get update && \
    apt-get install -y coreutils binutils build-essential libseccomp-dev gcc

RUN make -C ./nosocket/ all && make -C ./nosocket/ install

RUN bundle install
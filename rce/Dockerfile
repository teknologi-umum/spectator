FROM node:16.16.0-bullseye

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /app

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y coreutils binutils build-essential libseccomp-dev gcc apt-utils curl tar gzip

RUN curl -LO https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz && \
    tar -zxf grpcurl_1.8.6_linux_x86_64.tar.gz && \
    mv grpcurl /usr/bin/grpcurl && \
    chmod +x /usr/bin/grpcurl && \
    rm grpcurl_1.8.6_linux_x86_64.tar.gz LICENSE

COPY ./proto/rce.proto .

COPY ./rce/nosocket ./nosocket

RUN make -C ./nosocket/ all && make -C ./nosocket/ install

COPY ./rce/packages ./packages

COPY ./rce/scripts ./scripts

COPY ./rce/package*.json .

RUN npm install

RUN node ./scripts/install.cjs

RUN node ./scripts/register-users.cjs

COPY ./rce/tsconfig.json .

COPY ./rce/src ./src

COPY ./rce/build.js .

RUN npm run build

RUN rm -rf node_modules \
    && npm install --production

ENV PORT=50051

EXPOSE ${PORT}

CMD ["node", "./dist/index.js"]

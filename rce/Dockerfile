FROM node:16.14.0-bullseye

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /app

RUN apt-get update && \
    apt-get install -y coreutils binutils build-essential libseccomp-dev gcc apt-utils

COPY . .

RUN make -C ./nosocket/ all && make -C ./nosocket/ install

RUN npm install \
    && npm run build

RUN node ./scripts/install.cjs

RUN node ./scripts/register-users.cjs

RUN rm -rf node_modules \
    && npm install --production

ENV PORT=50051

EXPOSE ${PORT}

CMD ["node", "./dist/index.js"]

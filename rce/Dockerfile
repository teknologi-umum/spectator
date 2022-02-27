FROM node:16.14.0-bullseye

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /app

COPY . .

RUN apt-get update && \
    apt-get install -y coreutils binutils build-essential libseccomp-dev gcc apt-utils

RUN make -C ./nosocket/ all && make -C ./nosocket/ install

RUN npm install \
    && npm run build \
    && node ./scripts/install.cjs \
    && node ./scripts/register-users.cjs

RUN rm -rf node_modules \
    && npm install --production

CMD ["node", "./dist/index.js"]

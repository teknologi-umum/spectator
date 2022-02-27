FROM node:16.14.0-bullseye

WORKDIR /app

COPY . .

RUN apt-get update && \
    apt-get install -y coreutils binutils build-essential libseccomp-dev gcc

RUN make -C ./nosocket/ all && make -C ./nosocket/ install

RUN npm install \
    && npm run build \
    && node ./scripts/install.js \
    && node ./scripts/register-users.js

RUN rm -rf node_modules \
    && npm install --production

CMD ["node", "./dist/index.js"]

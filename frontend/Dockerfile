FROM node:16.13-alpine

WORKDIR /home/node/app

COPY . .

RUN npm install

RUN npm run build:prod \
    && rm -rf node_modules \
    && npm install --production

ENV NODE_ENV=production

EXPOSE 3000

CMD [ "node", "./node_modules/sirv-cli/bin.js", "--port", "3000", "./dist" ]

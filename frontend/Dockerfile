FROM node:16.13-alpine

WORKDIR /usr/src/app

COPY . .

RUN npm install

RUN npm run build:prod \
    && rm -rf node_modules \
    && npm install --production

ENV NODE_ENV=production

EXPOSE 3000

CMD [ "npm", "start" ]

FROM node:16.13-alpine

WORKDIR /home/node/app

COPY . .

RUN npm install

ENV VITE_SESSION_HUB_URL=/hubs/session
ENV VITE_EVENT_HUB_URL=/hubs/event
ENV VITE_LOGGER_URL=/log
ENV VITE_VIDEO_STREAM_URL=/video
ENV VITE_ADMIN_URL=/admin

RUN npm run build:prod \
    && rm -rf node_modules \
    && npm install --production

ENV NODE_ENV=production

EXPOSE 3000

CMD [ "node", "./node_modules/sirv-cli/bin.js", "--port", "3000", "./dist" ]

FROM node:12-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm ci

COPY . ./

ENV NODE_ENV=production

RUN npm run build

CMD npm run start

FROM node:12-alpine as builder
RUN apk update && apk add curl python make g++ && rm -rf /var/cache/apk/*
RUN curl -sfL https://install.goreleaser.com/github.com/tj/node-prune.sh | sh -s -- -b /usr/local/bin
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm ci
COPY . ./
ENV NODE_ENV=production
RUN npm run build
RUN npm prune
RUN /usr/local/bin/node-prune

FROM nginx:stable-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /usr/src/app/dist /bin/www
CMD [ "nginx", "-g", "daemon off;" ]

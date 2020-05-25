FROM node:12-alpine as builder
RUN apk update && apk add curl && rm -rf /var/cache/apk/*
RUN curl -sfL https://install.goreleaser.com/github.com/tj/node-prune.sh | sh -s -- -b /usr/local/bin
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm ci
COPY . ./
RUN npm run build
RUN npm prune --production
RUN /usr/local/bin/node-prune

FROM node:12-alpine
WORKDIR /usr/src/app
ENV NODE_ENV=production
COPY --from=builder /usr/src/app/.env ./
COPY --from=builder /usr/src/app/package.json ./
COPY --from=builder /usr/src/app/dist ./dist
COPY --from=builder /usr/src/app/src/graphql/schema ./dist/graphql/schema
COPY --from=builder /usr/src/app/node_modules ./node_modules
CMD [ "npm", "run", "start" ]

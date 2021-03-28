FROM node:14-alpine AS base
RUN apk add --no-cache g++ make python
WORKDIR /usr/src/app

FROM base AS dev
ENV NODE_ENV=development
CMD [ -d "node_modules" ] && npm run start || npm ci -q && npm run start

FROM base AS ci

FROM base AS prod-build
COPY package*.json ./
RUN npm ci -q
COPY . ./
ENV NODE_ENV=production
RUN npm run build

FROM lewislbr/gss:latest AS prod
COPY --from=prod-build /usr/src/app/gss.yaml ./
COPY --from=prod-build /usr/src/app/dist ./dist

FROM node:12-alpine as build
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm ci
COPY . ./
ENV NODE_ENV=production
RUN npm run build

FROM nginx:stable-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /usr/src/app/dist /bin/www
CMD [ "nginx", "-g", "daemon off;" ]

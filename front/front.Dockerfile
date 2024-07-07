FROM node:18 AS build

WORKDIR /app
COPY . .
RUN npm config set registry https://registry.npm.taobao.org/
RUN npm install --force
RUN npm run build

FROM nginx
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build /app/dist .
FROM node:18 AS build

WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

FROM nginx
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build /app/dist .
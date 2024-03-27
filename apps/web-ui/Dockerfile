# build stage
FROM node:lts-alpine as build-stage

RUN npm i -g http-server

WORKDIR /app

COPY package*.json ./
RUN npm i
RUN npm i vuetify

COPY . .
RUN npm run build

EXPOSE 8080
CMD ["http-server", "dist"]

# production stage
# FROM nginx:stable-alpine as prod-stage
# COPY --from=build-stage /app/dist /usr/share/nginx/html
#
# EXPOSE 80
#
# CMD ["nginx", "-g", "daemon off;"]

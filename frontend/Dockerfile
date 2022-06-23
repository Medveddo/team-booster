# NPM build stage
FROM node:18-alpine AS build

WORKDIR /usr/src/app
COPY package.json ./
COPY package-lock.json ./
RUN npm install

COPY . .

RUN npm run build

# Web server serves static
FROM caddy:2.4.6
COPY --from=build /usr/src/app/build /srv
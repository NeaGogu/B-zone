#Stage 1
FROM node:17-alpine as builder
WORKDIR /app
COPY frontend/package.json .
COPY frontend/package-lock.json .
RUN npm install
COPY frontend .
RUN npm run build

#Stage 2
FROM nginx:1.19.0
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/build .
ENTRYPOINT ["nginx", "-g", "daemon off;"]
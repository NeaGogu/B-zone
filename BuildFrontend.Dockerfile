#Stage 1
FROM node:17-alpine as builder
WORKDIR /app
COPY frontend/package.json .
COPY frontend/package-lock.json .
RUN npm install
COPY frontend .
RUN npm run build

#Stage 2
FROM nginx:1.21.3-alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/build .
COPY frontend_nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]
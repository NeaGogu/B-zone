# base image
FROM node:17-alpine

# set working directory
WORKDIR /app

# copy package.json and package-lock.json
COPY frontend/package*.json ./

# install dependencies
RUN npm install
COPY frontend .

# set environment variables
#ENV PORT=3000
ENV NODE_ENV=development

# expose port
EXPOSE $PORT

# start the application
CMD ["npm", "start"]

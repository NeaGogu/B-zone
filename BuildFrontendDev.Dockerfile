# base image
FROM node:14.17.6-alpine

# set working directory
WORKDIR /app

# copy package.json and package-lock.json
COPY package*.json ./

# install dependencies
RUN npm install

# copy the rest of the application code
COPY --from=builder /app/build .

# set environment variables
ENV PORT=3000
ENV NODE_ENV=development

# expose port
EXPOSE $PORT

# start the application
CMD ["npm", "start"]

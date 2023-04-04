# Use the official Golang image as the base image
FROM golang:1.20-alpine3.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY backend/ .

# Build the application inside the container
RUN go build -o b-zone ./cmd/web

ENV PORT=4000
ENV NODE_ENV=development

# Start the application
CMD ["./b-zone"]
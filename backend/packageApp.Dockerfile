# Use the official Golang image as the base image
FROM golang:1.20-alpine3.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application inside the container
RUN go build -o b-zone ./cmd/web

# Use a smaller base image for deployment
FROM alpine:3.16

# Copy the application binary from the build stage
COPY --from=0 /app/b-zone .

# Install any necessary system libraries or dependencies
# ca-certificates package contains a set of trusted root CA certificates that the app can use to verify SSL certificates.
RUN apk add --no-cache ca-certificates

# Expose the port that the application listens on
EXPOSE 4000

# Set any necessary environment variables
ENV MONGO_URL=mongodb://mongodb-container:27017
ENV PORT=4000

# Start the application
CMD ["./b-zone"]

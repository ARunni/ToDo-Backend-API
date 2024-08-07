# Use a base image with the correct Go version directly
FROM golang:1.22.3-alpine3.18 AS build-stage

# Set the working directory
WORKDIR /todo

# Copy the source code into the container
COPY ./ /todo

# Enable Go modules and download dependencies
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN mkdir -p /todo/build
RUN go mod download

# Build the application
RUN go build -v -o /todo/build/api ./cmd

# Use a minimal base image for the final stage
FROM gcr.io/distroless/static-debian11

# Copy the built binary and configuration file from the build stage
COPY --from=build-stage /todo/build/api /
COPY --from=build-stage /todo/.env /

# Expose the application port
EXPOSE 8000

# Run the application
CMD ["/api"]

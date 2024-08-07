# Stage 1: Build the Go application
FROM golang:1.22.3-alpine3.18 AS build-stage

# Set the working directory inside the container
WORKDIR /todo_svc

# Copy the entire project into the container
COPY ./ /todo_svc

# Create a build directory inside the container
RUN mkdir -p /todo_svc/build

# Download the Go module dependencies
RUN go mod download

# Build the Go application and place the output in the build directory
RUN go build -v -o /todo_svc/build/api ./cmd

# Stage 2: Create a minimal runtime image
FROM gcr.io/distroless/static-debian11

# Copy the built application and environment file from the build stage
COPY --from=build-stage /todo_svc/build/api /
COPY --from=build-stage /todo_svc/.env /

# Expose port 8080 for the application
EXPOSE 8080

# Set the command to run the application
CMD ["/api"]

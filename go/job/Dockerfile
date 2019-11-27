# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

#Set this environment var to use go modules
ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
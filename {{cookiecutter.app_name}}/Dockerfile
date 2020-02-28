
# Start from golang base image
FROM golang:1.13-alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY . ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go get -u 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o {{cookiecutter.app_name}} -mod vendor

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/{{cookiecutter.app_name}} .
# COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8800

#Command to run the executable
CMD ["./{{cookiecutter.app_name}} -hostport=0.0.0.0:8800"]
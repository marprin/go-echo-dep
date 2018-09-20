FROM golang:1.11

MAINTAINER Marprin <marprin93@gmail.com>

# Install support dependency
RUN apt update && apt install -y vim less

# Create the folder of working directory
RUN mkdir -p /app/src/package-manager

# Configure Golang
ENV GOPATH=/app

# # Include the package manager for golang
RUN go get -u github.com/golang/dep/cmd/dep

# Set alias for dep
RUN export dep='$GOPATH/bin/dep'

# Change working directory
WORKDIR /app/src/package-manager

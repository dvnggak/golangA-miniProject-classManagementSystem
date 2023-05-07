# syntax=docker/dockerfile:1 

FROM golang:1.17-alpine

# Create a working directory inside the image
WORKDIR /app

# Copy the modules and dependencies to image
COPY go.mod go.sum ./

# Download Go modules & dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./

# compile the Go app
RUN go build -o /miniproject

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["miniproject"]
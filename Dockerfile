# Use the official Golang image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o /godocker

# Command to run the Go app
CMD [ "/godocker" ]

# syntax=docker/dockerfile:1

FROM golang:1.23.3

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY *.go ./

# Build
RUN go build -o main main.go

EXPOSE 8080

# Run
CMD ["/main"]
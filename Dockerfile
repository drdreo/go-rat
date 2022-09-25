FROM golang:1.19 as builder
WORKDIR /usr/src/app

# Copy go mod and sum files 
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the working directory inside the container 
COPY . .
RUN go build -o bin/server app.go

EXPOSE 3000
CMD ["./bin/server"]
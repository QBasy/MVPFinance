# Use an official Go runtime as a parent image for building
FROM golang:1.20-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

COPY . .

RUN go build -o personal-finance-backend ./main.go

FROM ubuntu:latest

LABEL authors="Adilkhanov"

WORKDIR /root/

COPY --from=builder /app/personal-finance-backend .

EXPOSE 8080

ENTRYPOINT ["./personal-finance-backend"]

# stage 1 Build
FROM golang:1.25.0-alpine AS builder

# Sets /app as work dir inside the container
WORKDIR /app 

# copy mod and sum file to /app
COPY go.mod go.sum ./
# download dependencies 
RUN go mod download

# copy everything from local to destination (/app) inside our container
COPY . .

# compile go app into a binary (app)
RUN go build -o app ./cmd/api


# stage 2 Run
FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]

FROM golang:1.23.2 AS builder
WORKDIR /app

# not really needed for this project but just to make sure
COPY go.mod ./ 
RUN go mod download

# run go build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

# copying the binary to runner image
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .

# expose port and run
EXPOSE 8080
CMD ["./main"]

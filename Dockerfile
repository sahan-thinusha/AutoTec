FROM golang:1.18-alpine

RUN apk add --no-cache git

WORKDIR /

COPY . .
RUN apk add alpine-sdk

# Build the Go app

RUN go build -o app .

EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["./app"]
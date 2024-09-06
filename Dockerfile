FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go build ./cmd/api

EXPOSE 8081

# Run
CMD ["./api"]

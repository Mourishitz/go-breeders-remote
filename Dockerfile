FROM golang:latest

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build ./cmd/api

EXPOSE 8081

# Run
CMD ["./api"]

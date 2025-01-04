FROM golang:1.23-alpine

# Install necessary packages for CGO
RUN apk update && apk add --no-cache gcc musl-dev

# Set environment variables
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]
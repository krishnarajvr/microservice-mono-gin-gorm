# Development environment
# -----------------------
FROM golang:1.15-alpine
WORKDIR /micro

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate \
    && chmod +x /micro/docker/app/bin/*

EXPOSE 8080
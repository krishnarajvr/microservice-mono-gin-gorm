# Build environment
# -----------------
FROM golang:1.15-alpine as build-env
WORKDIR /micro

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate



# Deployment environment
# ----------------------
FROM alpine
RUN apk update

COPY --from=build-env /micro/bin/app /micro/
COPY --from=build-env /micro/bin/migrate /micro/
COPY --from=build-env /micro/migrations /micro/migrations

EXPOSE 8080
CMD ["/micro/app"]
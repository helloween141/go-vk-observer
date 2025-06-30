FROM golang:1.23-alpine AS build

RUN apk add --no-cache git

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/app/main.go

FROM alpine:3.19

RUN apk add --no-cache postgresql-client

COPY --from=build /main /app/main

COPY --from=build /go/bin/goose /usr/local/bin/goose

COPY --from=build /go/bin/sqlc /usr/local/bin/sqlc

EXPOSE 8080

COPY sqlc.yaml ./app

COPY ./internal/database/migrations /app/database/migrations

COPY ./docker/entrypoint.sh /app/entrypoint.sh
RUN chmod +x ./app/entrypoint.sh

ENTRYPOINT ["./app/entrypoint.sh"]
FROM golang:1.22.5-alpine AS build

# RUN apk add --no-cache bash curl \
#    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
#    && mv migrate.linux-amd64 /usr/local/bin/migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY ./internal/infrastructure/database/migrations ./migrations


FROM alpine:latest

COPY --from=build /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=build /app/migrations /migrations

RUN apk add --no-cache bash postgresql-client

ENTRYPOINT ["migrate"]

CMD ["-path", "/migrations", "-database", "${DATABASE_URI}", "up"]

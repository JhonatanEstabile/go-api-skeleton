FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN touch .env

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -o /api 

## Deploy
FROM scratch

WORKDIR /

COPY --from=build /api /api
COPY --from=build /app/.env .env
COPY --from=build /app/queries /queries

EXPOSE 8080

ENTRYPOINT ["/api"]
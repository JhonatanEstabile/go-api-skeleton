FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /api

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /api /api
COPY --from=build /app/queries /queries

RUN touch .env

EXPOSE 8080

ENTRYPOINT ["./api"]

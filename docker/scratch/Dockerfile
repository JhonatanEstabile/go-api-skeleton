##
## STEP 1 - BUILD
##
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /api

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
FROM clearlinux/golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go-app

EXPOSE 80

CMD [ "/go-app" ]

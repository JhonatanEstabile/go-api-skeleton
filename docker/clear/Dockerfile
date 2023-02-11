FROM clearlinux/golang:1.18

WORKDIR /app

COPY . .

RUN go mod download


RUN go build .

RUN touch .env

EXPOSE 8080

CMD [ "./api" ]

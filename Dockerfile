FROM golang

WORKDIR /usr/app

COPY . .

RUN go build

CMD ./employees-service


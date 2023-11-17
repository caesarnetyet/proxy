FROM golang:1.21.4 AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .docker.env .env

RUN go build -o app .

EXPOSE 3000

CMD [ "./app" ]
FROM golang:1.18.2

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build

EXPOSE 8080

CMD ["ls"]
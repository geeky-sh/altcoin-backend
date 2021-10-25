FROM golang:1.17-alpine

WORKDIR /app

COPY go.sum ./
COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-bitcoin

EXPOSE 8000

CMD ["/docker-gs-bitcoin"]
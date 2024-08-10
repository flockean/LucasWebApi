FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./lib ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /startApi

CMD [ "/startApi" ]

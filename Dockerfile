FROM golang:1.19.4

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY internal ./internal
COPY main.go ./

RUN go build -o /tictacgo

CMD ["/tictacgo"]

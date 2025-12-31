FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY paxos/ ./paxos/
COPY cmd/ ./cmd/

RUN go build -o paxos-server ./cmd/server/main.go

EXPOSE 8080

CMD ["./paxos-server"]

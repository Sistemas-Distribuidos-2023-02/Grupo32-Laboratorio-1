FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY central.go .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]
FROM golang:1.21

WORKDIR /central

COPY go.mod .
COPY central.go .
COPY parametros_de_inicio.txt .

RUN go get
RUN go build -o /docker-central

EXPOSE 8080

CMD ["/docker-central"]
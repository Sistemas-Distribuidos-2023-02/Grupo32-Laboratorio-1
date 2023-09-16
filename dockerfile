FROM golang:1.21

WORKDIR /central
COPY central.go .
RUN go mod init containerized-go-app
RUN go mod tidy

COPY parametros_de_inicio.txt .

RUN go get
RUN go build -o /docker-central

EXPOSE 8080

CMD ["/docker-central"]

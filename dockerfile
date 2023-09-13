FROM golang:1.21

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-central

EXPOSE 8080

CMD ["/docker-central"]
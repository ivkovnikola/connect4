FROM golang:1.14

WORKDIR /go/src/connect
COPY . .

RUN go install ./...

CMD ["connect"]
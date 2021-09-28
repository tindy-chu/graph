FROM golang:1.17.1-alpine3.14

WORKDIR /go/src/app

COPY ./app .

RUN go build -v -o app &&\
  mv app /go/bin

CMD tail -f /dev/null
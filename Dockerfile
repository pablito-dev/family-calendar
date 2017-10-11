FROM golang

ADD . /go/src/github.com/pablito-dev/family-calendar

RUN go get goji.io
RUN go get github.com/satori/go.uuid
RUN go get gopkg.in/mgo.v2

RUN go install github.com/pablito-dev/family-calendar

EXPOSE 8080
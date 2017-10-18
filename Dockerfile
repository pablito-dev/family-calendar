FROM golang

ADD . /go/src/github.com/pablito-dev/family-calendar

RUN go get goji.io \
  && go get github.com/satori/go.uuid \
  && go get gopkg.in/mgo.v2

RUN go install github.com/pablito-dev/family-calendar

ENTRYPOINT go/bin/family-calendar

EXPOSE 8080

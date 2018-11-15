FROM golang:1.11.2

RUN export GOPATH=$HOME/go
RUN export PATH=$GOPATH:$PATH

RUN adduser --disabled-password --gecos '' api
USER api

WORKDIR $GOPATH/src/github.com/kongebra/cpts/api

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2/bson

RUN go build -o main .

CMD ["./main"]
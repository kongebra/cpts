FROM golang:1.11.2

RUN adduser --disabled-password --gecos '' api
USER api

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2/bson

RUN mkdir /app
ADD . /app

RUN go build -o main .

CMD ["./main"]
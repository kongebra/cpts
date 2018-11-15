FROM golang:1.11.2

RUN adduser --disabled-password --gocos '' api
USER api

# RUN go get github.com/gorilla/mux
# RUN go get gopkg.in/mgo.v2/bson

RUN go get -d -v ./..
RUN go install -v ./..

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["./main"]
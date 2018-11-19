#FROM golang:onbuild
#EXPOSE 8080
FROM golang
WORKDIR /go/src/github.com/kongebra/cpts
COPY . .
RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=0 /go/src/github.com/kongebra/cpts .
EXPOSE 3000
CMD ["./cpts"]
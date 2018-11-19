FROM golang
WORKDIR /go/src/github.com/kongebra/cpts
COPY . .
RUN go get -d -v .

FROM scratch
WORKDIR /root/
COPY --from=0 /go/src/github.com/kongebra/cpts .
EXPOSE 3000
CMD ["./cpts"]
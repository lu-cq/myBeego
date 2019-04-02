FROM golang:latest
WORKDIR $GOPATH/src/github.com/myBeego
ADD . $GOPATH/src/github.com/myBeego
RUN go build .
EXPOSE 8080
ENTRYPOINT  ["./myBeego"]
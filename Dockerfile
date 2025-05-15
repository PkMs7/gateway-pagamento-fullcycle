FROM golang:1.23.9

WORKDIR /go/src

CMD [ "tail", "-f", "/dev/null" ]
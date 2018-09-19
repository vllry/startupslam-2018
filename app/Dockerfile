FROM golang:1.11
# https://hub.docker.com/_/golang/
# Note: it would be better to compile separately, and only ship the binary.

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

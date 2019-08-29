FROM golang:1.11 as builder
WORKDIR $GOPATH/src/github.com/johan-lejdung/go-microservice-api-guide/rest-api
COPY ./ .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
RUN cp rest-api /
RUN mkdir /db
RUN cp -r db/migrations/ /db/

FROM ubuntu:latest
WORKDIR /
COPY --from=builder /rest-api /
COPY --from=builder /db /db
CMD ["/rest-api"]

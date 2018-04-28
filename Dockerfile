FROM golang:latest AS build-env

RUN mkdir -p /go/src/github.com/lzientek/octopush-middleware
ADD . /go/src/github.com/lzientek/octopush-middleware/
WORKDIR /go/src/github.com/lzientek/octopush-middleware
RUN go build -o main .

FROM golang:alpine3.7
COPY --from=build-env /go/src/github.com/lzientek/octopush-middleware/main ./

EXPOSE 8080
ENTRYPOINT ["./main"]

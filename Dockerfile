# BUILD
FROM golang:1.11-alpine as builder

RUN apk add --no-cache git mercurial

ENV p $GOPATH/src/github.com/werberson/http-example

ADD ./ ${p}
WORKDIR ${p}
RUN go get -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /example .

# PKG
FROM scratch

COPY --from=builder /example /go/bin/

ENTRYPOINT [ "/go/bin/example" ]

EXPOSE 8080
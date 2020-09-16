FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/geordee/auths
COPY . .
RUN go get -d -v
RUN export CGO_ENABLED=0 && go build -o /go/bin/auths

FROM scratch

ENV DB_HOST=
ENV DB_USER=
ENV DB_PASS=
ENV DB_NAME=
ENV DB_PORT=

COPY --from=builder /go/bin/auths /go/bin/auths
ENTRYPOINT ["/go/bin/auths"]
EXPOSE 9090

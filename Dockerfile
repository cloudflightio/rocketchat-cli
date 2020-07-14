FROM golang:1.14-alpine as builder

LABEL maintainer="Michael Riedmann <michael.riedmann@cloudflight.io>"

WORKDIR $GOPATH/src/github.com/cloudflightio/rocketchat-cli

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -v ./...

FROM alpine:3.12

COPY --from=builder /go/bin/rocketchat-cli /bin/rocketchat-cli

ENTRYPOINT ["/bin/sh", "-c", "rocketchat-cli"]

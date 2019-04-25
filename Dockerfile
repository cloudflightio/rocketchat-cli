FROM golang:1.12 as builder

LABEL maintainer="Michael Riedmann <michael_riedmann@live.com>"

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR $GOPATH/src/github.com/mriedmann/rocketchat-cli

COPY . .

RUN dep ensure -vendor-only

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -v ./...

FROM scratch

COPY --from=builder /go/bin/rocketchat-cli /bin/rocketchat-cli

ENTRYPOINT ["/bin/rocketchat-cli"]

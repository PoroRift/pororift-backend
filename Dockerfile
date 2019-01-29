FROM golang:1.11
WORKDIR /go-source

RUN mkdir /go-source/bin
ENV GOBIN=/go-source/bin

COPY ./src ./src
WORKDIR /go-source/src
# RUN go test ./...
RUN go install ./...

ENTRYPOINT ["./bin/server"]
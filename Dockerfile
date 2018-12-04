FROM golang:1.11
WORKDIR /go-source

RUN mkdir /go-source/bin
ENV GOBIN=/go-source/bin

COPY ./src ./src
COPY ./go.mod .
COPY ./go.sum .

RUN go install ./src/server.go

ENTRYPOINT ["./bin/server"]
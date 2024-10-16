FROM gocv/opencv:4.8.1

ENV GOPATH /go

COPY . /go/src/cvpipe

WORKDIR /go/src/cvpipe 

RUN go test ./...

RUN go build -o main ./cmd/example/...

ENTRYPOINT ["/go/src/cvpipe/main"]

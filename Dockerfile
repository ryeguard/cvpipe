FROM gocv/opencv:latest

ENV GOPATH /go

COPY . /go/src/cvpipe

WORKDIR /go/src/cvpipe 
RUN go get -u -d gocv.io/x/gocv

RUN go build -o main ./cmd/example/...

ENTRYPOINT ["/go/src/cvpipe/main"]

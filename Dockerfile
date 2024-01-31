FROM gocv/opencv:latest

ENV GOPATH /go

COPY . /go/src/github.com/ryeguard/cvpipe

WORKDIR /go/src/github.com/ryeguard/cvpipe 
RUN go get -u -d gocv.io/x/gocv

RUN go build -o main ./cmd/main.go

ENTRYPOINT ["/go/src/github.com/ryeguard/cvpipe/main"]

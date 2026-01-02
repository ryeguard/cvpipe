FROM gocv/opencv:4.13.0

ENV GOPATH=/go

COPY . /go/src/cvpipe

WORKDIR /go/src/cvpipe 

RUN go test -v ./...

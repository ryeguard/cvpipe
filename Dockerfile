FROM ghcr.io/hybridgroup/opencv:4.10.0-static@sha256:3185e72e53b7df9ea12aab9334c69ec792758682fa71edb8a21323ba701b5d6d

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

COPY . /go/src/cvpipe

WORKDIR /go/src/cvpipe 

RUN ["go", "test", "./..."]
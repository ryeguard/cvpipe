# cvpipe

`cvpipe` is a tool for quickly prototyping computer vision filters and operations using [OpenCV](https://opencv.org/)/[GoCV](https://gocv.io/).

## Getting Started

Build and run the docker image to quickly get started:

```bash
docker build -t cvpipe .
```

```bash
docker run -it --rm \
--mount type=bind,source="$(pwd)/data",target=/go/src/cvpipe/data \
-t cvpipe
```

## Usage

The `cvpipe` package provides a simple API for chaining together OpenCV operations in Go. The following pseudo-code demonstrates the intention of the package:

```go
pipe := NewPipe()
defer pipe.Close()

result := pipe.
    Add(ResizeOperation(0.5)).
    Add(BlurOperation(5)).
    Add(ThresholdOperation(127)).
    Run(image)
defer result.Close()

// Do something with the result
```

For up-to-date, working/running examples, see the [cmd](./cmd) directory.

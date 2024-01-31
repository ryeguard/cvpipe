# cvpipe

`cvpipe` is a tool for quickly prototyping computer vision filters and operations using [OpenCV](https://opencv.org/)/[GoCV](https://gocv.io/).

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

See the [cmd](cmd) directory for running code examples.

package cvpipe

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestCloseNewPipe(t *testing.T) {
	pipe := NewPipe("test", PipeOptions{})
	defer pipe.Close()
}

func TestNewPipeMat(t *testing.T) {
	mat := gocv.NewMat()
	defer mat.Close()

	pipeMat := NewPipeMat(mat)
	defer pipeMat.Close()
}

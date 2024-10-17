package cvpipe

import (
	"image"
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

func TestRunEmptyPipe(t *testing.T) {
	pipe := NewPipe("test", PipeOptions{})
	defer pipe.Close()

	img := gocv.NewMat()
	defer img.Close()

	mat := NewPipeMat(img)
	defer mat.Close()

	result := pipe.Run(mat)
	defer result.Close()

	if !result.mat.Empty() {
		t.Error("non-empty result")
	}
}

func TestRunEmptyPipeWithBlankImage(t *testing.T) {
	pipe := NewPipe("test", PipeOptions{})
	defer pipe.Close()

	imageSize := 10

	img := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(0, 0, 0, 0), imageSize, imageSize, gocv.MatTypeCV8U)
	defer img.Close()

	mat := NewPipeMat(img)
	defer mat.Close()

	result := pipe.Run(mat)
	defer result.Close()

	if result.mat.Empty() {
		t.Error("empty result")
	}

	if result.mat.Rows() != imageSize || result.mat.Cols() != imageSize {
		t.Errorf("unexpected result size: got %dx%d, want %dx%d", result.mat.Rows(), result.mat.Cols(), imageSize, imageSize)
	}
}

func TestRunPipeInvertImage(t *testing.T) {
	pipe := NewPipe("test", PipeOptions{})
	defer pipe.Close()

	imageSize := 10

	img := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(255, 255, 255, 255), imageSize, imageSize, gocv.MatTypeCV8U)

	mat := NewPipeMat(img)
	defer mat.Close()

	resizeFactor := 0.5

	result := pipe.Add(NewResize(image.Point{}, resizeFactor, resizeFactor, gocv.InterpolationDefault)).Run(mat)
	defer result.Close()

	if result.mat.Empty() {
		t.Error("empty result")
	}

	wantSize := int(float64(imageSize) * resizeFactor)
	if result.mat.Rows() != wantSize || result.mat.Cols() != wantSize {
		t.Errorf("unexpected result size: got %dx%d, want %dx%d", result.mat.Rows(), result.mat.Cols(), wantSize, wantSize)
	}
}

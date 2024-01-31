package main

import (
	"image"
	"image/color"

	"github.com/ryeguard/cvpipe"
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("test.jpg", gocv.IMReadGrayScale)
	mat := cvpipe.NewPipeMat(img)
	defer mat.Close()

	pipe := cvpipe.NewPipe(cvpipe.PipeOptions{Save: true})
	defer pipe.Close()

	pipe.
		Add(cvpipe.NewResize(image.Pt(0, 0), 0.25, 0.25, gocv.InterpolationDefault), true).
		Add(cvpipe.NewGaussianBlur(image.Pt(15, 15), 0, 0, gocv.BorderDefault), true).
		Add(cvpipe.NewAdaptiveThreshold(255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 11, 2), true).
		Add(cvpipe.NewDilate(gocv.MorphRect, image.Pt(5, 5)), true).
		Add(cvpipe.NewPutText("Hello World!", image.Pt(10, 50), gocv.FontHersheyPlain, 1.0, color.RGBA{125, 125, 125, 125}, 2), true).
		Run(mat)
}

package main

import (
	"fmt"
	"image"

	"github.com/ryeguard/cvpipe"
	"github.com/ryeguard/cvpipe/colors"
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("./data/test.jpg", gocv.IMReadGrayScale)
	defer img.Close()
	if img.Empty() {
		fmt.Println("Error reading image")
		return
	}
	mat := cvpipe.NewPipeMat(img)
	defer mat.Close()

	pipe := cvpipe.NewPipe("example", cvpipe.PipeOptions{Save: true})
	defer pipe.Close()

	res := pipe.
		Add(cvpipe.NewResize(image.Pt(0, 0), 0.25, 0.25, gocv.InterpolationDefault)).
		Add(cvpipe.NewGaussianBlur(image.Pt(15, 15), 0, 0, gocv.BorderDefault)).
		Add(cvpipe.NewAdaptiveThreshold(255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 11, 2)).
		Add(cvpipe.NewDilate(gocv.MorphRect, image.Pt(5, 5))).
		Add(cvpipe.NewPutText("Hello World!", image.Pt(10, 50), gocv.FontHersheyPlain, 1.0, colors.Grey, 2)).
		Run(mat)
	defer res.Close()
}

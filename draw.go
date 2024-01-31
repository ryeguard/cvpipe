package cvpipe

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

type putText struct {
	text      string
	org       image.Point
	fontFace  gocv.HersheyFont
	fontScale float64
	color     color.RGBA
	thickness int
}

func (p *putText) Apply(m *PipeMat) *PipeMat {
	gocv.PutText(m.mat, p.text, p.org, p.fontFace, p.fontScale, p.color, p.thickness)
	return m
}

func (p *putText) Name() string {
	return "putText"
}

func NewPutText(text string, org image.Point, fontFace gocv.HersheyFont, fontScale float64, color color.RGBA, thickness int) Filter {
	return &putText{
		text:      text,
		org:       org,
		fontFace:  fontFace,
		fontScale: fontScale,
		color:     color,
		thickness: thickness,
	}
}

func (p *putText) Close() {
	return
}

type line struct {
	pt1       image.Point
	pt2       image.Point
	color     color.RGBA
	thickness int
}

func (l *line) Apply(m *PipeMat) *PipeMat {
	gocv.Line(m.mat, l.pt1, l.pt2, l.color, l.thickness)
	return m
}

func (l *line) Name() string {
	return "line"
}

func NewLine(pt1, pt2 image.Point, color color.RGBA, thickness int) Filter {
	return &line{
		pt1:       pt1,
		pt2:       pt2,
		color:     color,
		thickness: thickness,
	}
}

func (l *line) Close() {
	return
}

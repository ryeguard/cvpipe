package cvpipe

import (
	"image"

	"gocv.io/x/gocv"
)

type Filter interface {
	Name() string
	Apply(m *PipeMat) *PipeMat
	Close()
}

type threshold struct {
	thresh   float32
	maxValue float32
	typ      gocv.ThresholdType
}

func NewThreshold(thresh, maxValue float32, typ gocv.ThresholdType) Filter {
	return &threshold{
		thresh:   thresh,
		maxValue: maxValue,
		typ:      typ,
	}
}

func (t *threshold) Apply(m *PipeMat) *PipeMat {
	gocv.Threshold(*m.mat, m.temp, t.thresh, t.maxValue, t.typ)
	m.mat = m.temp
	return m
}

func (t *threshold) Name() string {
	return "threshold"
}

func (t *threshold) Close() {
	return
}

type gaussianBlur struct {
	kSize  image.Point
	sigmaX float64
	sigmaY float64
	typ    gocv.BorderType
}

func NewGaussianBlur(ksize image.Point, sigmaX, sigmaY float64, typ gocv.BorderType) Filter {
	return &gaussianBlur{
		kSize:  ksize,
		sigmaX: sigmaX,
		sigmaY: sigmaY,
		typ:    typ,
	}
}

func (g *gaussianBlur) Apply(m *PipeMat) *PipeMat {
	gocv.GaussianBlur(*m.mat, m.temp, g.kSize, g.sigmaX, g.sigmaY, g.typ)
	m.mat = m.temp
	return m
}

func (g *gaussianBlur) Name() string {
	return "gaussianBlur"
}

func (g *gaussianBlur) Close() {
	return
}

type canny struct {
	threshold1 float32
	threshold2 float32
}

func NewCanny(threshold1, threshold2 float32) Filter {
	return &canny{
		threshold1: threshold1,
		threshold2: threshold2,
	}
}

func (c *canny) Apply(m *PipeMat) *PipeMat {
	gocv.Canny(*m.mat, m.temp, c.threshold1, c.threshold2)
	m.mat = m.temp
	return m
}

func (c *canny) Name() string {
	return "canny"
}

func (c *canny) Close() {
	return
}

type sobel struct {
	ddepth gocv.MatType
	dx     int
	dy     int
	ksize  int
	scale  float64
	delta  float64
	typ    gocv.BorderType
}

func NewSobel(ddepth gocv.MatType, dx, dy, ksize int, scale, delta float64, typ gocv.BorderType) Filter {
	return &sobel{
		dx:    dx,
		dy:    dy,
		ksize: ksize,
		scale: scale,
		delta: delta,
		typ:   typ,
	}
}

func (s *sobel) Apply(m *PipeMat) *PipeMat {
	gocv.Sobel(*m.mat, m.temp, m.mat.Type(), s.dx, s.dy, s.ksize, s.scale, s.delta, s.typ)
	m.mat = m.temp
	return m
}

func (s *sobel) Name() string {
	return "sobel"
}

func (s *sobel) Close() {
	return
}

type adaptiveThreshold struct {
	maxValue    float32
	adaptiveTyp gocv.AdaptiveThresholdType
	typ         gocv.ThresholdType
	blockSize   int
	c           float32
}

func (a *adaptiveThreshold) Apply(m *PipeMat) *PipeMat {
	gocv.AdaptiveThreshold(*m.mat, m.temp, a.maxValue, a.adaptiveTyp, a.typ, a.blockSize, a.c)
	m.mat = m.temp
	return m
}

func (a *adaptiveThreshold) Name() string {
	return "adaptiveThreshold"
}

func NewAdaptiveThreshold(maxValue float32, adaptiveTyp gocv.AdaptiveThresholdType, typ gocv.ThresholdType, blockSize int, c float32) Filter {
	return &adaptiveThreshold{
		maxValue:    maxValue,
		adaptiveTyp: adaptiveTyp,
		typ:         typ,
		blockSize:   blockSize,
		c:           c,
	}
}

func (a *adaptiveThreshold) Close() {
	return
}

type dilate struct {
	kernel gocv.Mat
}

func (d *dilate) Apply(m *PipeMat) *PipeMat {
	gocv.Dilate(*m.mat, m.temp, d.kernel)
	m.mat = m.temp
	return m
}

func (d *dilate) Name() string {
	return "dilate"
}

func NewDilate(shape gocv.MorphShape, ksize image.Point) Filter {
	kernel := gocv.GetStructuringElement(shape, ksize)
	return &dilate{
		kernel: kernel,
	}
}

func (d *dilate) Close() {
	d.kernel.Close()
}

type erode struct {
	kernel gocv.Mat
}

func (e *erode) Apply(m *PipeMat) *PipeMat {
	gocv.Erode(*m.mat, m.temp, e.kernel)
	m.mat = m.temp
	return m
}

func (e *erode) Name() string {
	return "erode"
}

func NewErode(shape gocv.MorphShape, ksize image.Point) Filter {
	kernel := gocv.GetStructuringElement(shape, ksize)
	return &erode{
		kernel: kernel,
	}
}

func (e *erode) Close() {
	e.kernel.Close()
}

type resize struct {
	sz     image.Point
	fx     float64
	fy     float64
	interp gocv.InterpolationFlags
}

func (r *resize) Apply(m *PipeMat) *PipeMat {
	gocv.Resize(*m.mat, m.temp, r.sz, r.fx, r.fy, r.interp)
	m.mat = m.temp
	return m
}

func (r *resize) Name() string {
	return "resize"
}

func NewResize(dsize image.Point, fx, fy float64, typ gocv.InterpolationFlags) Filter {
	return &resize{
		sz:     dsize,
		fx:     fx,
		fy:     fy,
		interp: typ,
	}
}

func (r *resize) Close() {
	return
}

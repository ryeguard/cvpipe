package cvpipe

import (
	"fmt"

	"gocv.io/x/gocv"
)

type PipeOptions struct {
	// Enable saving of intermediate results. Will be overridden by individual step save options.
	Save bool
}

type PipeMat struct {
	mat    *gocv.Mat
	temp   *gocv.Mat
	kernel *gocv.Mat
}

type Pipe struct {
	steps   []step
	opts    PipeOptions
	results []*PipeMat
}

type step struct {
	Filter
	save bool
}

func NewPipe(opts PipeOptions) *Pipe {
	return &Pipe{opts: opts}
}

func (p *Pipe) Close() {
	for _, s := range p.steps {
		s.Close()
	}
}

// Add a filter to the pipe. If save is true, the result of this step will be saved to disk.
func (p *Pipe) Add(f Filter, save bool) *Pipe {
	p.steps = append(p.steps, step{f, save})
	return p
}

func (p *Pipe) Run(m *PipeMat) *PipeMat {
	for i, f := range p.steps {
		m = f.Apply(m)
		if f.save || (f.save && p.opts.Save) {
			gocv.IMWrite(fmt.Sprintf("data/%03d_%v.jpg", i, f.Name()), *m.mat)
		}
	}
	return m
}

func NewPipeMat(mat gocv.Mat) *PipeMat {
	temp := gocv.NewMat()
	kernel := gocv.NewMat()
	return &PipeMat{
		mat:    &mat,
		temp:   &temp,
		kernel: &kernel,
	}
}

func (m *PipeMat) Close() {
	m.mat.Close()
	m.temp.Close()
	m.kernel.Close()
}

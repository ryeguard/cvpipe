package cvpipe

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

type PipeOptions struct {
	// Enable saving of intermediate results. Will be overridden by individual step save options.
	Save bool
}

type FilterOptions struct {
	// Enable saving of this step's result. Takes precedence over Pipe options
	Save *bool
}

type PipeMat struct {
	mat    *gocv.Mat
	temp   *gocv.Mat
	kernel *gocv.Mat
}

type Pipe struct {
	name    string
	steps   []step
	opts    PipeOptions
	results []*PipeMat
}

type step struct {
	Filter
	opts FilterOptions
}

func NewPipe(name string, opts PipeOptions) *Pipe {
	return &Pipe{
		name: name,
		opts: opts,
	}
}

func (p *Pipe) Close() {
	for _, s := range p.steps {
		s.Close()
	}
}

// Add a filter to the pipe. If save is true, the result of this step will be saved to disk.
func (p *Pipe) Add(f Filter, opts ...FilterOptions) *Pipe {
	if len(opts) == 0 {
		p.steps = append(p.steps, step{f, FilterOptions{}})
		return p
	}

	p.steps = append(p.steps, step{f, opts[0]})
	return p
}

// Run applies all filters in the pipe to the given PipeMat.
func (p *Pipe) Run(m *PipeMat) *PipeMat {
	for i, f := range p.steps {
		m = f.Apply(m)

		var shouldSave bool
		if f.opts.Save == nil {
			shouldSave = p.opts.Save
		} else {
			shouldSave = *f.opts.Save
		}

		if shouldSave {
			gocv.IMWrite(fmt.Sprintf("data/%v_%03d_%v.jpg", p.name, i, f.Name()), *m.mat)
		}
	}
	return m
}

type MatOptions struct {
	Size *image.Point
	Type *gocv.MatType
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

func NewPipeMatFromOptions(opts MatOptions) (*PipeMat, error) {
	if opts.Size == nil && opts.Type == nil {
		return nil, fmt.Errorf("Size and Type must be set")
	}

	mat := gocv.NewMatWithSize(opts.Size.Y, opts.Size.X, *opts.Type)
	temp := gocv.NewMat()
	kernel := gocv.NewMat()

	return &PipeMat{
		mat:    &mat,
		temp:   &temp,
		kernel: &kernel,
	}, nil
}

func (m *PipeMat) Close() {
	m.mat.Close()
	m.temp.Close()
	m.kernel.Close()
}

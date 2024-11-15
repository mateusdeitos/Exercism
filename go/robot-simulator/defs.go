package robot

import "fmt"

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'

func (c *Command) IsValid() bool {
	return *c == 'R' || *c == 'L' || *c == 'A'
}

type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

func (r *Step2Robot) TurnLeft() {
	r.Dir--
	if r.Dir < 0 {
		r.Dir = W
	}
}

func (r *Step2Robot) TurnRight() {
	r.Dir++
	if r.Dir > W {
		r.Dir = N
	}
}

func (r *Step2Robot) Advance() {
	switch r.Dir {
	case N:
		r.Northing++
	case E:
		r.Easting++
	case S:
		r.Northing--
	case W:
		r.Easting--
	}
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}

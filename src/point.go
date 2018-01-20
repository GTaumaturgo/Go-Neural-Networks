package main

import (
	"math/rand"
	"time"
)

type point struct {
	x     float64
	y     float64
	label int
}

func newPoint() point {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	p := point{
		x: (r.Float64() * 20) - 10,
		y: (r.Float64() * 20) - 10,
	}

	if p.x > p.y {
		p.label = -1
	} else {
		p.label = 1
	}

	return p
}

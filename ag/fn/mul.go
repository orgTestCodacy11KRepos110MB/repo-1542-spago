// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"sync"

	"github.com/nlpodyssey/spago/mat"
)

// Mul is an operator to perform matrix-vector multiplication.
type Mul[O Operand] struct {
	x1 O // matrix
	x2 O // vector
}

// NewMul returns a new Mul Function.
func NewMul[O Operand](x1 O, x2 O) *Mul[O] {
	return &Mul[O]{
		x1: x1,
		x2: x2,
	}
}

// Operands returns the list of operands.
func (r *Mul[O]) Operands() []O {
	return []O{r.x1, r.x2}
}

// Forward computes the output of the function.
func (r *Mul[O]) Forward() mat.Matrix {
	return r.x1.Value().Mul(r.x2.Value())
}

// Backward computes the backward pass.
func (r *Mul[O]) Backward(gy mat.Matrix) {
	if !(r.x1.Value().Rows() == gy.Rows() && r.x2.Value().Columns() == gy.Columns()) {
		panic("fn: matrices with not compatible size")
	}
	var wg sync.WaitGroup
	if r.x1.RequiresGrad() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x2t := r.x2.Value().T()
			defer mat.ReleaseMatrix(x2t)
			gx := gy.Mul(x2t)
			defer mat.ReleaseMatrix(gx)
			r.x1.AccGrad(gx)
		}()
	}
	if r.x2.RequiresGrad() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//r.x2.AccGrad(gy.T().Mul(r.x1).T()) // alternative method
			if gy.Columns() == 1 {
				gx := r.x1.Value().MulT(gy)
				defer mat.ReleaseMatrix(gx)
				r.x2.AccGrad(gx)
			} else {
				x1t := r.x1.Value().T()
				defer mat.ReleaseMatrix(x1t)
				gx := x1t.Mul(gy)
				defer mat.ReleaseMatrix(gx)
				r.x2.AccGrad(gx)
			}
		}()
	}
	wg.Wait()
}

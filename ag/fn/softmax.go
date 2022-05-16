// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
)

// Softmax is a single-input softmax function.
type Softmax[O Operand] struct {
	x        O
	y        mat.Matrix // initialized during the forward pass (required by the backward pass)
	operands []O
}

// NewSoftmax returns a new Softmax Function.
func NewSoftmax[O Operand](x O) *Softmax[O] {
	return &Softmax[O]{
		x:        x,
		operands: []O{x},
	}
}

// Operands returns the list of operands.
func (r *Softmax[O]) Operands() []O {
	return r.operands
}

// Forward computes the output of this function.
func (r *Softmax[O]) Forward() mat.Matrix {
	r.y = r.x.Value().Softmax()
	return r.y
}

// Backward computes the backward pass.
func (r *Softmax[O]) Backward(gy mat.Matrix) {
	if !(mat.SameDims(r.x.Value(), gy) || mat.VectorsOfSameSize(r.x.Value(), gy)) {
		panic("fn: matrices with not compatible size")
	}
	if r.x.RequiresGrad() {
		y := r.y
		n := y.Size()
		jb := y.NewInitFuncMatrix(n, n, func(row, col int) float.Float {
			// FIXME: avoid casting to specific type
			vRow := y.ScalarAtVec(row).F64()
			if row == col {
				return float.Interface(vRow * (1 - vRow))
			}
			vCol := y.ScalarAtVec(col).F64()
			return float.Interface(-(vRow * vCol))
		})
		defer mat.ReleaseMatrix(jb)
		gx := jb.Mul(gy)
		defer mat.ReleaseMatrix(gx)
		r.x.AccGrad(gx)
	}
}

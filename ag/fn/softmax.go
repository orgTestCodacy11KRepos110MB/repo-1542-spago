// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"github.com/nlpodyssey/spago/mat"
)

var _ Function[float32] = &Softmax[float32]{}

// Softmax is a single-input softmax function.
type Softmax[T mat.DType] struct {
	x Operand[T]
	y mat.Matrix[T] // initialized during the forward pass (required by the backward pass)
}

// NewSoftmax returns a new Softmax Function.
func NewSoftmax[T mat.DType](x Operand[T]) *Softmax[T] {
	return &Softmax[T]{x: x}
}

// Forward computes the output of this function.
func (r *Softmax[T]) Forward() mat.Matrix[T] {
	r.y = r.x.Value().Softmax()
	return r.y
}

// Backward computes the backward pass.
func (r *Softmax[T]) Backward(gy mat.Matrix[T]) {
	if !(mat.SameDims(r.x.Value(), gy) || mat.VectorsOfSameSize(r.x.Value(), gy)) {
		panic("fn: matrices with not compatible size")
	}
	if r.x.RequiresGrad() {
		y := r.y
		n := y.Size()
		jb := mat.NewInitFuncDense[T](n, n, func(row, col int) T {
			if row == col {
				v := y.AtVec(row)
				return v * (1 - v)
			}
			return -(y.AtVec(row) * y.AtVec(col))
		})
		defer mat.ReleaseDense(jb)
		gx := jb.Mul(gy)
		defer mat.ReleaseMatrix(gx)
		r.x.PropagateGrad(gx)
	}
}

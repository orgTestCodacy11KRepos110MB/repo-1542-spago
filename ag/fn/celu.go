// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"github.com/nlpodyssey/spago/mat"
)

var _ Function[float32] = &CELU[float32]{}

// CELU is an operator to perform the CELU activation.
// CELU(x) = max(0,x) + min(0,α ∗ (exp(x/α) − 1))
type CELU[T mat.DType] struct {
	x     Operand[T]
	alpha Operand[T] // scalar
}

// NewCELU returns a new CELU Function.
func NewCELU[T mat.DType](x, alpha Operand[T]) *CELU[T] {
	return &CELU[T]{x: x, alpha: alpha}
}

// Forward computes the output of the function.
func (r *CELU[T]) Forward() mat.Matrix[T] {
	return r.x.Value().ApplyWithAlpha(celu[T], r.alpha.Value().Scalar())
}

// Backward computes the backward pass.
func (r *CELU[T]) Backward(gy mat.Matrix[T]) {
	if !(mat.SameDims(r.x.Value(), gy) || mat.VectorsOfSameSize(r.x.Value(), gy)) {
		panic("fn: matrices with not compatible size")
	}
	if r.x.RequiresGrad() {
		gx := r.x.Value().ApplyWithAlpha(celuDeriv[T], r.alpha.Value().Scalar())
		defer mat.ReleaseMatrix(gx)
		gx.ProdInPlace(gy)
		r.x.PropagateGrad(gx)
	}
}
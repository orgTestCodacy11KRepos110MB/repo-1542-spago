// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import "github.com/nlpodyssey/spago/mat"

// Stack is a Function which stacks together all given operand matrices,
// producing a single bigger matrix as result.
type Stack[O Operand] struct {
	xs []O
}

// NewStack returns a new Stack Function.
func NewStack[O Operand](xs []O) *Stack[O] {
	return &Stack[O]{xs: xs}
}

// Operands returns the list of operands.
func (r *Stack[O]) Operands() []O {
	return r.xs
}

// Forward computes the output of the function.
func (r *Stack[O]) Forward() mat.Matrix {
	if len(r.xs) == 0 {
		panic("fn: Stack has no operands")
	}
	vs := make([]mat.Matrix, len(r.xs))
	for i, x := range r.xs {
		vs[i] = x.Value()
	}
	return vs[0].NewStack(vs...)
}

// Backward computes the backward pass.
func (r *Stack[O]) Backward(gy mat.Matrix) {
	if gy.Rows() != len(r.xs) {
		panic("fn: matrices with not compatible size")
	}

	for i, x := range r.xs {
		if !x.RequiresGrad() {
			continue
		}
		gyRow := gy.ExtractRow(i).ReshapeInPlace(x.Value().Dims())
		x.AccGrad(gyRow)
		mat.ReleaseMatrix(gyRow)
	}
}

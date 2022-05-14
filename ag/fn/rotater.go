// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import "github.com/nlpodyssey/spago/mat"

// RotateR is a function to perform a right circular shift of a vector.
type RotateR[T mat.DType, O Operand[T]] struct {
	x        O
	i        int
	operands []O
}

// NewRotateR returns a new RotateR Function. `i` is the number of places by
// which the elements are shifted.
func NewRotateR[T mat.DType, O Operand[T]](x O, i int) *RotateR[T, O] {
	return &RotateR[T, O]{
		x:        x,
		i:        i,
		operands: []O{x},
	}
}

// Operands returns the list of operands.
func (r *RotateR[T, O]) Operands() []O {
	return r.operands
}

// Forward computes the output of the function.
func (r *RotateR[T, O]) Forward() mat.Matrix[T] {
	return mat.NewVecDense(rotateR(mat.Data[T](r.x.Value()), r.i))
}

// Backward computes the backward pass.
func (r *RotateR[T, O]) Backward(gy mat.Matrix[T]) {
	if r.x.RequiresGrad() {
		gx := mat.NewVecDense(rotateL(mat.Data[T](gy), r.i))
		defer mat.ReleaseDense(gx)
		r.x.AccGrad(gx)
	}
}

func rotateR[T mat.DType](a []T, i int) []T {
	x, b := a[:(len(a)-i)], a[(len(a)-i):]
	return append(b, x...)
}

func rotateL[T mat.DType](a []T, i int) []T {
	x, b := a[:i], a[i:]
	return append(b, x...)
}

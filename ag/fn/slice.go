// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import "github.com/nlpodyssey/spago/mat"

// Slice is a function to extract a portion of a matrix.
type Slice[O Operand] struct {
	x       O
	fromRow int
	fromCol int
	toRow   int
	toCol   int
}

// NewSlice returns a new Slice Function.
func NewSlice[O Operand](x O, fromRow, fromCol, toRow, toCol int) *Slice[O] {
	return &Slice[O]{
		x:       x,
		fromRow: fromRow,
		fromCol: fromCol,
		toRow:   toRow,
		toCol:   toCol,
	}
}

// Operands returns the list of operands.
func (s *Slice[O]) Operands() []O {
	return []O{s.x}
}

// Forward computes the output of the function.
func (s *Slice[O]) Forward() mat.Matrix {
	return s.x.Value().Slice(s.fromRow, s.fromCol, s.toRow, s.toCol)
}

// Backward computes the backward pass.
func (s *Slice[O]) Backward(gy mat.Matrix) {
	lx := s.toRow - s.fromRow
	ly := s.toCol - s.fromCol
	if !(gy.Rows() == lx && gy.Columns() == ly) {
		panic("fn: matrices with not compatible size")
	}
	if s.x.RequiresGrad() {
		gx := s.x.Value().ZerosLike()
		defer mat.ReleaseMatrix(gx)
		for i := 0; i < lx; i++ {
			for j := 0; j < ly; j++ {
				gx.SetScalar(i+s.fromRow, j+s.fromCol, gy.ScalarAt(i, j))
			}
		}
		s.x.AccGrad(gx)
	}
}

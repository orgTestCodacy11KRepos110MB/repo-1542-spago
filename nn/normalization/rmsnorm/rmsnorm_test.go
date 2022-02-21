// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rmsnorm

import (
	"github.com/nlpodyssey/spago/ag"
	"github.com/nlpodyssey/spago/mat"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel_Forward(t *testing.T) {
	t.Run("float32", testModelForward[float32])
	t.Run("float64", testModelForward[float64])
}

func testModelForward[T mat.DType](t *testing.T) {
	model := newTestModel[T]()
	g := ag.NewGraph[T](ag.WithMode[T](ag.Training))
	// == Forward
	x1 := g.NewVariable(mat.NewVecDense([]T{1.0, 2.0, 0.0, 4.0}), true)
	x2 := g.NewVariable(mat.NewVecDense([]T{3.0, 2.0, 1.0, 6.0}), true)
	x3 := g.NewVariable(mat.NewVecDense([]T{6.0, 2.0, 5.0, 1.0}), true)

	y := ag.Bind(g, model).Forward(x1, x2, x3)

	assert.InDeltaSlice(t, []T{0.6182178902, 0.1254256878, 0.2, 1.4965944974}, y[0].Value().Data(), 1.0e-06)
	assert.InDeltaSlice(t, []T{0.8242640687, 0.186862915, 0.2848528137, 1.4576450198}, y[1].Value().Data(), 1.0e-06)
	assert.InDeltaSlice(t, []T{1.1385489459, 0.2015268072, 0.5692744729, 0.2969463856}, y[2].Value().Data(), 1.0e-06)

	// == Backward
	y[0].PropagateGrad(mat.NewVecDense([]T{-1.0, -0.2, 0.4, 0.6}))
	y[1].PropagateGrad(mat.NewVecDense([]T{-0.3, 0.1, 0.7, 0.9}))
	y[2].PropagateGrad(mat.NewVecDense([]T{0.3, -0.4, 0.7, -0.8}))
	g.BackwardAll()

	assert.InDeltaSlice(t, []T{-0.2493918746, -0.0448905374, 0.0523722937, 0.0847932373}, x1.Grad().Data(), 1.0e-06)
	assert.InDeltaSlice(t, []T{-0.1109874804, -0.0513642366, 0.0365432785, 0.066524606}, x2.Grad().Data(), 1.0e-06)
	assert.InDeltaSlice(t, []T{0.0040284488, 0.0087283057, 0.0242825941, -0.1630402749}, x3.Grad().Data(), 1.0e-06)
}

func newTestModel[T mat.DType]() *Model[T] {
	model := New[T](4)
	model.W.Value().SetData([]T{0.5, -0.2, 0.3, 0.8})
	model.B.Value().SetData([]T{0.4, 0.3, 0.2, 0.1})
	return model
}

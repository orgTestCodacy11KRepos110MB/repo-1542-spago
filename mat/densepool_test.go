// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/nlpodyssey/spago/mat/float"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDensePool(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		dp := densePool[float32]()
		require.Same(t, densePoolFloat32, dp)
	})

	t.Run("float64", func(t *testing.T) {
		dp := densePool[float64]()
		require.Same(t, densePoolFloat64, dp)
	})
}

func TestDensePool_GetDense(t *testing.T) {
	t.Run("float32", testDensePoolGetDense[float32])
	t.Run("float64", testDensePoolGetDense[float64])
}

func testDensePoolGetDense[T float.DType](t *testing.T) {
	densePool := densePool[T]()

	t.Run("negative rows", func(t *testing.T) {
		require.Panics(t, func() {
			densePool.Get(-1, 1)
		})
	})

	t.Run("negative cols", func(t *testing.T) {
		require.Panics(t, func() {
			densePool.Get(1, -1)
		})
	})

	testCases := []struct {
		rows        int
		cols        int
		expectedCap int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},

		{1, 1, 1},

		{1, 2, 3},
		{2, 1, 3},

		{1, 3, 3},
		{3, 1, 3},

		{1, 4, 7},
		{4, 1, 7},

		{5, 5, 31},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d x %d", tc.rows, tc.cols), func(t *testing.T) {
			d := densePool.Get(tc.rows, tc.cols)
			assertDenseFromPoolDims(t, tc.rows, tc.cols, tc.expectedCap, d)
		})
	}
}

func assertDenseFromPoolDims[T float.DType](t *testing.T, expectedRows, expectedCols, expectedCap int, d *Dense[T]) {
	t.Helper()
	expectedSize := expectedRows * expectedCols
	assert.Equal(t, expectedRows, d.rows)
	assert.Equal(t, expectedCols, d.cols)
	assert.Len(t, d.data, expectedSize)
	assert.Equal(t, expectedCap, cap(d.data))
}

func TestDensePool_GetAndRelease(t *testing.T) {
	t.Run("float32", testGetAndRelease[float32])
	t.Run("float64", testGetAndRelease[float64])
}

func testGetAndRelease[T float.DType](t *testing.T) {
	runtime.GC()
	a1 := densePool[T]().Get(5, 1)
	b1 := densePool[T]().Get(10, 1)

	assert.Len(t, a1.data, 5)
	assert.Equal(t, 7, cap(a1.data))

	assert.Len(t, b1.data, 10)
	assert.Equal(t, 15, cap(b1.data))

	a1.data[0] = 42
	b1.data[0] = 24

	densePool[T]().Put(a1)
	densePool[T]().Put(b1)

	a2 := densePool[T]().Get(6, 1)
	b2 := densePool[T]().Get(9, 1)

	x := densePool[T]().Get(6, 1)
	y := densePool[T]().Get(9, 1)

	assert.Len(t, a2.data, 6)
	assert.Equal(t, 7, cap(a2.data))

	assert.Len(t, b2.data, 9)
	assert.Equal(t, 15, cap(b2.data))

	assert.Len(t, x.data, 6)
	assert.Equal(t, 7, cap(x.data))

	assert.Len(t, y.data, 9)
	assert.Equal(t, 15, cap(y.data))

	if a2.data[0] != 42 {
		t.Errorf("a1 and a2 should share the same slice data")
	}
	if b2.data[0] != 24 {
		t.Errorf("b1 and b2 should share the same slice data")
	}
	if x.data[0] != 0 {
		t.Errorf("slice data of `x` should be blank")
	}
	if y.data[0] != 0 {
		t.Errorf("slice data of `y` should be blank")
	}
	runtime.GC()
}

func TestDensePool_Get(t *testing.T) {
	t.Run("float32", testDensePoolGet[float32])
	t.Run("float64", testDensePoolGet[float64])
}

func testDensePoolGet[T float.DType](t *testing.T) {
	runtime.GC()
	d := densePool[T]().Get(2, 3)

	assert.Equal(t, 2, d.Rows())
	assert.Equal(t, 3, d.Columns())
	assert.Equal(t, []T{0, 0, 0, 0, 0, 0}, Data[T](d))

	SetData[T](d, []T{1, 2, 3, 4, 5, 6})
	densePool[T]().Put(d)
	d = densePool[T]().Get(2, 3)
	assert.Equal(t, []T{1, 2, 3, 4, 5, 6}, Data[T](d), "possible dirty data is not zeroed")
	densePool[T]().Put(d)
	runtime.GC()
}

func TestDensePool_GetEmpty(t *testing.T) {
	t.Run("float32", testDensePoolGetEmpty[float32])
	t.Run("float64", testDensePoolGetEmpty[float64])
}

func testDensePoolGetEmpty[T float.DType](t *testing.T) {
	d := densePool[T]().GetEmpty(2, 3)

	assert.Equal(t, 2, d.Rows())
	assert.Equal(t, 3, d.Columns())
	assert.Equal(t, []T{0, 0, 0, 0, 0, 0}, Data[T](d))

	SetData[T](d, []T{1, 2, 3, 4, 5, 6})
	densePool[T]().Put(d)
	d = densePool[T]().GetEmpty(2, 3)
	assert.Equal(t, []T{0, 0, 0, 0, 0, 0}, Data[T](d), "possible dirty data is zeroed")
	densePool[T]().Put(d)
}

func TestDensePool_Put(t *testing.T) {
	t.Run("float32", testDensePoolPut[float32])
	t.Run("float64", testDensePoolPut[float64])
}

func testDensePoolPut[T float.DType](t *testing.T) {
	t.Run("it panics if the matrix does not come from the workspace", func(t *testing.T) {
		d := NewEmptyDense[T](3, 4)
		view := d.View(4, 3)
		assert.Panics(t, func() { densePool[T]().Put(view.(*Dense[T])) })
		densePool[T]().Put(d)
	})
}

func TestReleaseDense(t *testing.T) {
	t.Run("float32", testReleaseDense[float32])
	t.Run("float64", testReleaseDense[float64])
}

func testReleaseDense[T float.DType](t *testing.T) {
	t.Run("it panics if the matrix does not come from the workspace", func(t *testing.T) {
		d := NewEmptyDense[T](3, 4)
		view := d.View(4, 3)
		assert.Panics(t, func() { ReleaseDense(view.(*Dense[T])) })
		ReleaseDense(d)
	})
}

func TestReleaseMatrix(t *testing.T) {
	t.Run("float32", testReleaseMatrix[float32])
	t.Run("float64", testReleaseMatrix[float64])
}

func testReleaseMatrix[T float.DType](t *testing.T) {
	t.Run("it panics if the matrix does not come from the workspace", func(t *testing.T) {
		d := NewEmptyDense[T](3, 4)
		view := d.View(4, 3)
		assert.Panics(t, func() { ReleaseMatrix(view) })
		ReleaseMatrix(d)
	})

	t.Run("it panics if the matrix is not of Dense", func(t *testing.T) {
		d := new(foreignMatrixImplementation[T])
		assert.Panics(t, func() { ReleaseMatrix(d) })
	})
}

type foreignMatrixImplementation[T float.DType] struct {
	*Dense[T]
}

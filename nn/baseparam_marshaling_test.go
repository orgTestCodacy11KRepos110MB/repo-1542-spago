// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nn

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/mat/mattest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParam_Gob(t *testing.T) {
	t.Run("float32", testParamGob[float32])
	t.Run("float64", testParamGob[float64])
}

func testParamGob[T float.DType](t *testing.T) {
	t.Run("with all serializable values set", func(t *testing.T) {
		var buf bytes.Buffer

		p1 := NewParam(mat.NewScalar(T(12)))
		p1.SetName("foo")
		p1.SetType(Biases)
		p1.SetRequiresGrad(false)
		p1.SetPayload(&Payload{
			Label: 42,
			Data:  []mat.Matrix{mat.NewScalar(T(34))},
		})

		err := gob.NewEncoder(&buf).Encode(p1)
		require.Nil(t, err)

		var p2 *BaseParam

		err = gob.NewDecoder(&buf).Decode(&p2)
		require.Nil(t, err)
		require.NotNil(t, p2)
		require.NotNil(t, p2.Value())
		mattest.AssertMatrixEquals(t, mat.NewScalar(T(12)), p2.Value())
		assert.Equal(t, "foo", p2.Name())
		assert.Equal(t, Biases, p2.Type())
		assert.False(t, p2.RequiresGrad())

		payload := p2.Payload()
		assert.NotNil(t, payload)
		assert.Equal(t, 42, payload.Label)
		assert.NotEmpty(t, payload.Data)
		mattest.AssertMatrixEquals(t, mat.NewScalar(T(34)), payload.Data[0])
	})

	t.Run("with default properties and nil value", func(t *testing.T) {
		var buf bytes.Buffer

		p1 := NewParam(nil)

		err := gob.NewEncoder(&buf).Encode(p1)
		require.Nil(t, err)

		var p2 *BaseParam
		err = gob.NewDecoder(&buf).Decode(&p2)
		require.Nil(t, err)
		require.NotNil(t, p2)
		assert.Nil(t, p2.Value())
		assert.Empty(t, p2.Name())
		assert.Equal(t, Undefined, p2.Type())
		assert.True(t, p2.RequiresGrad())
		assert.Nil(t, p2.Payload())
	})
}

// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot_test

import (
	"bytes"
	"testing"

	"github.com/nlpodyssey/spago/ag"
	"github.com/nlpodyssey/spago/ag/encoding"
	"github.com/nlpodyssey/spago/ag/encoding/dot"
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	t.Run("float32", testEncode[float32])
	t.Run("float64", testEncode[float64])
}

func testEncode[T float.DType](t *testing.T) {
	t.Run("without time steps", func(t *testing.T) {
		a := ag.Var(mat.NewScalar[T](1)).WithGrad(false).WithName("a")
		b := ag.Var(mat.NewScalar[T](3)).WithGrad(false).WithName("b")
		y := ag.Sum(a, b)

		g := encoding.NewGraph(y)
		buf := new(bytes.Buffer)
		err := dot.Encode(g, buf)
		require.NoError(t, err)
		out := buf.String()

		expectedOutput := `strict digraph {
	rankdir=LR;
    ordering="in";
    outputMode="edgesfirst";
	colorscheme="dark28";
	node [colorscheme="dark28"];

	1->0;
	2->0;

	0 [label=<<sup>0</sup><br/><b>Add</b><br/><sub>1×1</sub>>,shape=oval]
	1 [label=<<sup>1</sup><br/><b>a</b><br/><sub>1×1</sub>>,shape=box]
	2 [label=<<sup>2</sup><br/><b>b</b><br/><sub>1×1</sub>>,shape=box]
}
`
		assert.Equal(t, expectedOutput, out)
	})

	t.Run("with time steps", func(t *testing.T) {
		tsh := ag.NewTimeStepHandler()

		a := ag.Var(mat.NewScalar[T](1)).WithGrad(false).WithName("a")
		b := ag.Var(mat.NewScalar[T](3)).WithGrad(false).WithName("b")
		c := ag.Var(mat.NewScalar[T](5)).WithGrad(false).WithName("c")
		d := ag.Var(mat.NewScalar[T](7)).WithGrad(false).WithName("d")

		x := ag.Sum(a, b)
		tsh.IncTimeStep()
		y := ag.Sum(x, c)
		z := ag.Sum(y, d)

		g := encoding.NewGraph(z).WithTimeSteps(tsh)
		buf := new(bytes.Buffer)
		err := dot.Encode(g, buf)
		require.NoError(t, err)
		out := buf.String()

		expectedOutput := `strict digraph {
	rankdir=LR;
    ordering="in";
    outputMode="edgesfirst";
	colorscheme="dark28";
	node [colorscheme="dark28"];

	1->0;
	2->1;
	3->2;
	4->2;
	5->1;
	6->0;

	subgraph cluster_timestep_0 {
		label="Time Step 0";
		color=1;
		node [color=1];
		
		2 [label=<<sup>2</sup><br/><b>Add</b><br/><sub>1×1</sub>>,shape=oval]
		3 [label=<<sup>3</sup><br/><b>a</b><br/><sub>1×1</sub>>,shape=box]
		4 [label=<<sup>4</sup><br/><b>b</b><br/><sub>1×1</sub>>,shape=box]
		5 [label=<<sup>5</sup><br/><b>c</b><br/><sub>1×1</sub>>,shape=box]
		6 [label=<<sup>6</sup><br/><b>d</b><br/><sub>1×1</sub>>,shape=box]
	}

	subgraph cluster_timestep_1 {
		label="Time Step 1";
		color=2;
		node [color=2];
		
		0 [label=<<sup>0</sup><br/><b>Add</b><br/><sub>1×1</sub>>,shape=oval]
		1 [label=<<sup>1</sup><br/><b>Add</b><br/><sub>1×1</sub>>,shape=oval]
	}

}
`
		assert.Equal(t, expectedOutput, out)
	})
}

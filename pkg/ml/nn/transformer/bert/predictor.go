// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bert

import (
	"github.com/nlpodyssey/spago/pkg/ml/ag"
	"github.com/nlpodyssey/spago/pkg/ml/nn/activation"
	"github.com/nlpodyssey/spago/pkg/ml/nn/linear"
	"github.com/nlpodyssey/spago/pkg/ml/nn/normalization/layernorm"
	"github.com/nlpodyssey/spago/pkg/ml/nn/stack"
)

type PredictorConfig struct {
	InputSize        int
	HiddenSize       int
	OutputSize       int
	HiddenActivation ag.OpName
	OutputActivation ag.OpName
}

func NewPredictor(config PredictorConfig) *stack.Model {
	return stack.New(
		linear.New(config.InputSize, config.HiddenSize),
		activation.New(config.HiddenActivation),
		layernorm.New(config.HiddenSize),
		linear.New(config.HiddenSize, config.OutputSize),
		activation.New(config.OutputActivation),
	)
}

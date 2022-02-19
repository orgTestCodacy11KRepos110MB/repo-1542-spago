// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlpmixer implements the MLP-Mixer (Tolstikhin et al., 2021).
package mlpmixer

import (
	"encoding/gob"
	"fmt"

	"github.com/nlpodyssey/spago/ag"
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
	"github.com/nlpodyssey/spago/nn/normalization/layernorm"
)

var _ nn.Model = &MixerBlock[float32]{}

// MixerBlock contains the serializable parameters.
type MixerBlock[T mat.DType] struct {
	nn.BaseModel
	Config
	TokenLayerNorm   *layernorm.Model[T]
	TokenMixerFF     *FeedForward[T]
	ChannelLayerNorm *layernorm.Model[T]
	ChannelMixerFF   *FeedForward[T]
}

// Config provides configuration settings for a MixerBlock.
type Config struct {
	InputSize               int
	HiddenSizeTokenMixer    int
	HiddenSizeChannelMixer  int
	Channels                int
	ActFunctionTokenMixer   activation.Name
	ActFunctionChannelMixer activation.Name
}

func init() {
	gob.Register(&MixerBlock[float32]{})
	gob.Register(&MixerBlock[float64]{})
}

// New returns a new model with parameters initialized to zeros.
func New[T mat.DType](config Config) *MixerBlock[T] {
	return &MixerBlock[T]{
		Config:           config,
		TokenMixerFF:     newFeedForward[T](config.Channels, config.HiddenSizeTokenMixer, config.ActFunctionTokenMixer, 0),
		TokenLayerNorm:   layernorm.New[T](config.InputSize),
		ChannelMixerFF:   newFeedForward[T](config.InputSize, config.HiddenSizeChannelMixer, config.ActFunctionChannelMixer, 0),
		ChannelLayerNorm: layernorm.New[T](config.InputSize),
	}
}

// Forward performs the forward step for each input node and returns the result.
func (m *MixerBlock[T]) Forward(xs ...ag.Node[T]) []ag.Node[T] {
	if len(xs) > m.Config.Channels {
		panic(fmt.Sprintf("mlpmixer: maximum sequence length is %d, got %d",
			m.Config.Channels, len(xs)))
	}

	xs = m.residual(m.tokenMix(xs), xs)
	xs = m.residual(m.channelMix(xs), xs)
	return xs
}

func (m *MixerBlock[T]) tokenMix(xs []ag.Node[T]) ag.Node[T] {
	normalized := m.TokenLayerNorm.Forward(xs...)
	stacked := ag.Stack(normalized...)

	ys := make([]ag.Node[T], stacked.Value().Columns())
	for i := range ys {
		col := ag.ColView(stacked, i)
		ys[i] = m.TokenMixerFF.Forward(col)[0]
	}
	return ag.T(ag.Stack(ys...))
}

func (m *MixerBlock[T]) channelMix(xs []ag.Node[T]) ag.Node[T] {
	normalized := m.ChannelLayerNorm.Forward(xs...)

	ys := make([]ag.Node[T], len(normalized))
	for i, nv := range normalized {
		nvt := ag.T(nv)
		ys[i] = m.ChannelMixerFF.Forward(nvt)[0]
	}
	return ag.Stack(ys...)
}

func (m *MixerBlock[T]) residual(x ag.Node[T], residual []ag.Node[T]) []ag.Node[T] {
	ys := make([]ag.Node[T], x.Value().Rows())
	for i := 0; i < x.Value().Rows(); i++ {
		ys[i] = ag.Add(ag.RowView(x, i), residual[i])
	}
	return ys
}

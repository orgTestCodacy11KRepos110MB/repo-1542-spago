// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nn

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/nlpodyssey/spago/mat"
)

// init registers the param implementation with the gob subsystem - so that it knows how to encode and decode
// values of type nn.Param
func init() {
	gob.Register(&BaseParam{})
}

type baseParamForMarshaling struct {
	Name         string
	PType        ParamsType
	Value        mat.Matrix
	Payload      *Payload
	RequiresGrad bool
}

// MarshalBinary marshals a param into binary form.
func (p *BaseParam) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	v := baseParamForMarshaling{
		Name:         p.name,
		PType:        p.pType,
		Value:        p.value,
		Payload:      p.payload,
		RequiresGrad: p.requiresGrad,
	}
	err := enc.Encode(v)
	if err != nil {
		return nil, fmt.Errorf("cannot encode BaseParam: %w", err)
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary unmarshals a param from binary form.
func (p *BaseParam) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)
	dec := gob.NewDecoder(r)
	var v baseParamForMarshaling
	err := dec.Decode(&v)
	if err != nil {
		return fmt.Errorf("cannot decode BaseParam: %w", err)
	}
	p.name = v.Name
	p.pType = v.PType
	p.value = v.Value
	p.payload = v.Payload
	p.requiresGrad = v.RequiresGrad
	return nil
}

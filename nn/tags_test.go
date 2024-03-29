// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidModuleFieldTagParsing(t *testing.T) {
	for _, example := range []struct {
		tag      string
		expected moduleFieldTag
	}{
		{"", moduleFieldTag{
			Type: defaultModuleFieldType,
		}},
		{"type:weights", moduleFieldTag{
			Type: weightsModuleFieldType,
		}},
		{"type:biases", moduleFieldTag{
			Type: biasesModuleFieldType,
		}},
		{"type:undefined", moduleFieldTag{
			Type: undefinedModuleFieldType,
		}},
	} {
		t.Run(fmt.Sprintf("%#v", example.tag), func(t *testing.T) {
			actual, err := parseModuleFieldTag(example.tag)
			assert.Nil(t, err)
			assert.Equal(t, example.expected, actual)
		})
	}
}

func TestInvalidModuleFieldTagParsing(t *testing.T) {
	for _, example := range []string{
		" ",
		"foo",
		"foo:bar",
		"type:foo",
	} {
		t.Run(fmt.Sprintf("%#v", example), func(t *testing.T) {
			_, err := parseModuleFieldTag(example)
			assert.NotNil(t, err)
		})
	}
}

func TestModuleFieldTag_ParamType(t *testing.T) {
	for _, example := range []struct {
		tag      string
		expected ParamsType
	}{
		{"type:weights", Weights},
		{"type:biases", Biases},
		{"type:undefined", Undefined},
	} {
		t.Run(fmt.Sprintf("%#v", example.tag), func(t *testing.T) {
			mft, err := parseModuleFieldTag(example.tag)
			assert.Nil(t, err)
			assert.Equal(t, example.expected, mft.paramType())
		})
	}
}

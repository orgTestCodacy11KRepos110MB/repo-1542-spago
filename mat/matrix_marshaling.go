// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	binaryMatrixNil byte = iota
	binaryMatrixDense32
	binaryMatrixDense64
)

// MarshalBinaryMatrix encodes a Matrix into binary form.
func MarshalBinaryMatrix(m Matrix, w io.Writer) error {
	var identifier byte
	var data []byte
	var err error

	switch mt := m.(type) {
	case nil:
		_, err = w.Write([]byte{binaryMatrixNil})
		return err
	case *Dense[float32]:
		identifier = binaryMatrixDense32
		data, err = mt.MarshalBinary()
		if err != nil {
			return err
		}
	case *Dense[float64]:
		identifier = binaryMatrixDense64
		data, err = mt.MarshalBinary()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("mat: unexpected matrix type %T", mt)
	}

	idAndSize := [9]byte{}
	idAndSize[0] = identifier
	binary.LittleEndian.PutUint64(idAndSize[1:], uint64(len(data)))

	_, err = w.Write(idAndSize[:])
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// UnmarshalBinaryMatrix decodes a Matrix from binary form.
func UnmarshalBinaryMatrix(r io.Reader) (Matrix, error) {
	idAndSize := [9]byte{}

	_, err := r.Read(idAndSize[:1])
	if err != nil {
		return nil, err
	}

	identifier := idAndSize[0]
	if identifier == binaryMatrixNil {
		return nil, nil
	}

	_, err = r.Read(idAndSize[1:])
	if err != nil {
		return nil, err
	}
	dataSize := int(binary.LittleEndian.Uint64(idAndSize[1:]))

	data := make([]byte, dataSize)
	_, err = r.Read(data)
	if err != nil {
		return nil, err
	}

	switch identifier {
	case binaryMatrixDense32:
		d := new(Dense[float32])
		err = d.UnmarshalBinary(data)
		if err != nil {
			return nil, err
		}
		return d, nil
	case binaryMatrixDense64:
		d := new(Dense[float64])
		err = d.UnmarshalBinary(data)
		if err != nil {
			return nil, err
		}
		return d, nil
	default:
		return nil, fmt.Errorf("mat: unexpected matrix identifier %d", identifier)
	}
}

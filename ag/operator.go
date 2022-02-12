// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ag

import (
	"fmt"
	"github.com/nlpodyssey/spago/ag/fn"
	"github.com/nlpodyssey/spago/mat"
	"reflect"
	"sync"
)

var (
	_ fn.Operand[float32] = &Operator[float32]{}
	_ GradValue[float32]  = &Operator[float32]{}
	_ Node[float32]       = &Operator[float32]{}
)

var (
	operatorPoolFloat32 = &sync.Pool{
		New: func() interface{} { return new(Operator[float32]) },
	}
	operatorPoolFloat64 = &sync.Pool{
		New: func() interface{} { return new(Operator[float64]) },
	}
)

func getOperatorPool[T mat.DType]() *sync.Pool {
	// TODO: review this code once stable go 1.18 is released
	switch any(T(0)).(type) {
	case float32:
		return any(operatorPoolFloat32).(*sync.Pool)
	case float64:
		return any(operatorPoolFloat64).(*sync.Pool)
	default:
		panic(fmt.Sprintf("ag: no operator pool for type %T", T(0)))
	}
}

// Operator is a type of node.
type Operator[T mat.DType] struct {
	graph        *Graph[T]
	timeStep     int
	id           int
	function     fn.Function[T]
	operands     []Node[T]
	value        mat.Matrix[T] // store the results of a forward evaluation
	mu           sync.Mutex    // to avoid data race during gradients accumulation
	grad         mat.Matrix[T] // TODO: support of sparse gradients
	hasGrad      bool
	requiresGrad bool
}

// ID returns the ID of the node in the graph.
func (r *Operator[_]) ID() int {
	return r.id
}

// Name returns the Name of the operator.
// The name is taken from the name of r.function via/ reflection.
func (r *Operator[_]) Name() string {
	return reflect.ValueOf(r.function).Elem().Type().Name()
}

// Graph returns the graph this node belongs to.
func (r *Operator[T]) Graph() *Graph[T] {
	return r.graph
}

// Value returns the cached result of the function.
func (r *Operator[T]) Value() mat.Matrix[T] {
	return r.value
}

// ScalarValue returns the the scalar value of the node.
// It panics if the value is not a scalar.
// Note that it is not possible to start the backward step from a scalar value.
func (r *Operator[T]) ScalarValue() T {
	return r.value.Scalar()
}

// Grad returns the gradients accumulated during the backward pass.
func (r *Operator[T]) Grad() mat.Matrix[T] {
	return r.grad
}

// PropagateGrad accumulates the gradients to the node itself.
func (r *Operator[T]) PropagateGrad(grad mat.Matrix[T]) {
	if !r.requiresGrad {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.grad == nil {
		r.grad = r.value.ZerosLike()
	}
	r.grad.AddInPlace(grad)
	r.hasGrad = true
}

// HasGrad returns true if there are accumulated gradients.
func (r *Operator[_]) HasGrad() bool {
	return r.hasGrad
}

// RequiresGrad returns true if the node requires gradients.
func (r *Operator[_]) RequiresGrad() bool {
	return r.requiresGrad
}

// ZeroGrad clears the gradients.
func (r *Operator[_]) ZeroGrad() {
	if r.grad == nil {
		return
	}
	defer mat.ReleaseMatrix(r.grad) // release memory
	r.grad = nil
	r.hasGrad = false
}

// TimeStep returns the time-step of the node.
func (r *Operator[_]) TimeStep() int {
	return r.timeStep
}

// Operands returns the operands of the operator.
func (r *Operator[T]) Operands() []Node[T] {
	return r.operands
}

func (r *Operator[_]) backward() {
	if !r.hasGrad {
		return
	}
	r.function.Backward(r.grad)
}
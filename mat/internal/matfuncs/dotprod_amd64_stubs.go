// Code generated by command: go run dotprod_asm.go -out ../../matfuncs/dotprod_amd64.s -stubs ../../matfuncs/dotprod_amd64_stubs.go -pkg matfuncs. DO NOT EDIT.

//go:build amd64 && gc && !purego

package matfuncs

// DotProdAVX32 returns the dot product between x1 and x2 (32 bits, AVX required).
//
//go:noescape
func DotProdAVX32(x1 []float32, x2 []float32) float32

// DotProdAVX64 returns the dot product between x1 and x2 (64 bits, AVX required).
//
//go:noescape
func DotProdAVX64(x1 []float64, x2 []float64) float64

// DotProdSSE32 returns the dot product between x1 and x2 (32 bits, SSE required).
//
//go:noescape
func DotProdSSE32(x1 []float32, x2 []float32) float32

// DotProdSSE64 returns the dot product between x1 and x2 (64 bits, SSE required).
//
//go:noescape
func DotProdSSE64(x1 []float64, x2 []float64) float64

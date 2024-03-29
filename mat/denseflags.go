// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

type denseFlag byte

const (
	denseIsFromPool denseFlag = 1 << iota
	denseIsNew
	denseIsView
)

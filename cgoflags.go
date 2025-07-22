// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5go

// #cgo CFLAGS: -I${SRCDIR}/../../../../third_party/hdf5/include
// #cgo LDFLAGS: -L${SRCDIR}/../../../../third_party/hdf5/lib -lhdf5_hl -lhdf5
// #include "hdf5.h"
import "C"

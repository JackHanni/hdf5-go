// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	hdf5 "github.com/JackHanni/hdf5-go"
)

func main() {
	fmt.Println("=== go-hdf5 ===")
	version, err := hdf5.LibVersion()
	if err != nil {
		fmt.Printf("** error ** %s\n", err)
		return
	}
	fmt.Printf("=== version: %s", version)
	fmt.Println("=== bye.")
}

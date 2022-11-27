// Copyright 2022 The Go Authors in ByteDance Inc. All rights reserved.
// Use of this source code should follow terms of ByteDance Inc.
//
// Authors:     hujinghui.jeffrey@bytedance.com
// Create Date: 2022-11-27
// Update Date: 2022-11-27
//

package main

import (
	"fmt"
)

func main() {
	total := 100
	sum := 0
	for i := 0; i < total; i++ {
		sum += i
	}
	fmt.Println(sum)
}

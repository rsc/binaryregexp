// +build go1.10

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

import "strings"

// Buffer is a type alias to strings.Builder to allow
// this package to be usable on >=Go1.10.
// See https://code-review.googlesource.com/c/gocloud/+/41410
type Buffer = strings.Builder

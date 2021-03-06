// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"testing"

	"go.etcd.io/etcd/v3/pkg/testutil"
)

func TestMain(m *testing.M) {
	testutil.MustTestMainWithLeakDetection(m)
}

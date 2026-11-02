// Copyright (c) 2026 The Monetarium developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpc

import (
	"testing"

	mondtypes "github.com/monetarium/monetarium-node/rpc/jsonrpc/types"
)

func TestCheckVersionNodeRenumbering(t *testing.T) {
	tests := []struct {
		major, minor, patch uint32
		ok                  bool
	}{
		{1, 3, 14, true}, // v1.3.13+ release line
		{1, 3, 0, true},
		{2, 1, 0, true}, // before the renumbering
		{1, 2, 9, false},
		{2, 0, 5, false},
		{3, 0, 0, false},
	}
	for _, tc := range tests {
		verMap := map[string]mondtypes.VersionResult{
			"monetarium": {Major: tc.major, Minor: tc.minor, Patch: tc.patch},
		}
		err := checkVersion(verMap, "monetarium")
		if (err == nil) != tc.ok {
			t.Errorf("%d.%d.%d: err=%v, want ok=%v", tc.major, tc.minor, tc.patch, err, tc.ok)
		}
	}
}

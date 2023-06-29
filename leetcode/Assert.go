package leetcode

import "testing"

func Assert(t *testing.T, assert any, result any) {
	if assert != result {
		t.Errorf("Error, %s, %s", assert, result)
	}
}

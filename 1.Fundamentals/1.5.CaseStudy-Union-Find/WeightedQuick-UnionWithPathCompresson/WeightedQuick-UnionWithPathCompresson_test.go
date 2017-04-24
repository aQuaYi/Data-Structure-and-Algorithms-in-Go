package wquwpc

import (
	"testing"
)

var u = &uf{}

func Test_New(t *testing.T) {
	n := 100
	u = New(n)
	if u.count != n {
		t.Error("")
	}
}

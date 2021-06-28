package gostubTest

import (
	"github.com/prashantv/gostub"
	"testing"
)

var counter = 100

func TestStubGlobalVariable(t *testing.T) {
	stubs := gostub.Stub(&counter, 200)
	defer stubs.Reset()
	t.Log("Counter:", counter)
}

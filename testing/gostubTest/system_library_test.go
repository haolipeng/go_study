package gostubTest

import (
	"github.com/prashantv/gostub"
	"os"
	"testing"
	"time"
)

var timeNow = time.Now //not time.Now

func GetDate() int {
	return timeNow().Year()
}

func TestStubSystemLibrary(t *testing.T) {
	t.Log("before stub,GetDate():", GetDate())
	stubs := gostub.StubFunc(&timeNow, time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC))
	t.Log("after stub,GetDate():", GetDate())
	defer stubs.Reset()
}

var osHostName = os.Hostname

func GetHostName() string {
	var host string
	var err error
	if host, err = osHostName(); err != nil {
		return ""
	}
	return host
}

func TestStubMultipleValueSystemLibrary(t *testing.T) {
	t.Log("before stub,GetHostName():", GetHostName())
	stubs := gostub.StubFunc(&osHostName, "haolipeng", nil)
	defer stubs.Reset()
	t.Log("after stub,GetHostName():", GetHostName())
}

package framework

import (
	"testing"
)

func TestMD5Hash(t *testing.T) {
	if MD5Hash("12345678") != "25d55ad283aa400af464c76d713c07ad" {
		t.Error("Wrong generated hash")
	}
}

package even

import "testing"

func TestEven(t *testing.T) {
	if true != Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}

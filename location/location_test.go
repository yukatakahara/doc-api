package location

import "testing"

func TestLocation(t *testing.T) {
	expected := "foo"
	actual := Sort()
	if actual != expected {
		t.Error("Test failed")
	}
}

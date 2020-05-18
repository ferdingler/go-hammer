package core

import "testing"

func TestUuid(t *testing.T) {
	one := uuid()
	two := uuid()
	if one == two {
		t.Errorf("Expected uuids to be different, %s, %s", one, two)
	}
}

package excav

import "testing"

func TestJoinParams(t *testing.T) {
	p1 := Params{
		"KEY1": "P1 VALUE1",
		"KEY2": "P1 VALUE2",
		"KEY3": "P1 VALUE3",
	}

	p2 := Params{
		"KEY2": "P2 VALUE2",
		"KEY4": "P2 VALUE4",
	}

	p3 := MergeParams(p1, p2)

	if p3["KEY4"] != "P2 VALUE4" {
		t.FailNow()
	}

	if p3["KEY2"] != "P2 VALUE2" {
		t.FailNow()
	}

}

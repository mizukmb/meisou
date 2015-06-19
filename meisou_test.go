package timer

import (
	"testing"
)

func TestCanUseTimer(t *testing.T) {
	num := [5]int{-1, 0, 1, 59, 61}
	actual, _ := canUseTimer(num[0])
	expect := false
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}

	actual, _ = canUseTimer(num[1])
	expect = false
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}

	actual, _ = canUseTimer(num[2])
	expect = true
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}

	actual, _ = canUseTimer(num[3])
	expect = true
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}

	actual, _ = canUseTimer(num[4])
	expect = false
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}
}

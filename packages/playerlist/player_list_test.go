package playerlist

import "testing"

func TestRemove(t *testing.T) {
	p, err := Create("playerlist")

	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	l := len(p)

	p.Remove("Andrew Wiggins")

	if len(p) >= l {
		t.Errorf("Entry not removed %d, %d, %v", l, len(p), p)
		return
	}

	if p.Contains("Andrew Wiggins") {
		t.Errorf("Entry not removed %v", p)
		return
	}
}

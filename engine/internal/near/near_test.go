package near

import "testing"

func TestClosest(t *testing.T) {
	cands := []string{"Clamp", "Abs", "Noise", "YClampedGradient"}
	for in, want := range map[string]string{
		"Clmp":   "Clamp",
		"clamp":  "Clamp", // case only
		"Nosie":  "Noise",
		"Zzzzzz": "",
		"":       "",
	} {
		if got := Closest(in, cands); got != want {
			t.Errorf("Closest(%q) = %q, want %q", in, got, want)
		}
	}
	// Ties go to the name that sorts first, whatever the order.
	if a, b := Closest("ab", []string{"ac", "aa"}), Closest("ab", []string{"aa", "ac"}); a != "aa" || b != "aa" {
		t.Errorf("tie: %q, %q", a, b)
	}
	if got := ClosestKey("hal_entry", map[string]bool{"hall_entry": true, "room_entry": true}); got != "hall_entry" {
		t.Errorf("ClosestKey = %q", got)
	}
	if d := Distance("kitten", "sitting"); d != 3 {
		t.Errorf("Distance = %d", d)
	}
}

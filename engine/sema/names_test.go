package sema

import "testing"

func TestCase(t *testing.T) {
	for _, c := range []struct{ pascal, snake string }{
		{"XzScale", "xz_scale"}, {"YAbove", "y_above"}, {"AboveBottom", "above_bottom"},
		{"Noise", "noise"}, {"OldBlendedNoise", "old_blended_noise"}, {"LootTable", "loot_table"},
	} {
		if got := Snake(c.pascal); got != c.snake {
			t.Errorf("Snake(%s) = %s, want %s", c.pascal, got, c.snake)
		}
		if got := Pascal(c.snake); got != c.pascal {
			t.Errorf("Pascal(%s) = %s, want %s", c.snake, got, c.pascal)
		}
	}
}

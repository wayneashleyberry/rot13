package rot13

import "testing"

func TestIndexPosition(t *testing.T) {
	var pos int
	pos = indexPosition("y", []string{"x", "y", "z"})
	if pos != 1 {
		t.Errorf("should be 1, got %v", pos)
	}

	pos = indexPosition("z", []string{"x", "y", "z"})
	if pos != 2 {
		t.Errorf("should be 2, got %v", pos)
	}
}

func TestEncode(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{
			"Why did the chicken cross the road?", "Jul qvq gur puvpxra pebff gur ebnq?",
		},
		{
			"!@#$", "!@#$",
		},
	}

	for _, c := range cases {
		got := Encode(c.in)
		if got != c.want {
			t.Errorf("Expected: %q Got: %q Using %q", c.want, got, c.in)
		}
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{
			"Gb trg gb gur bgure fvqr!", "To get to the other side!",
		},
		{
			"...", "...",
		},
	}

	for _, c := range cases {
		got := Decode(c.in)
		if got != c.want {
			t.Errorf("Expected: %q Got: %q Using %q", c.want, got, c.in)
		}
	}
}

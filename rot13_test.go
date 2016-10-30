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
	input := "Why did the chicken cross the road?"
	expected := "Jul qvq gur puvpxra pebff gur ebnq?"
	output := Encode(input)

	if output != expected {
		t.Errorf("Expected: %q Got: %q Using %q", expected, output, input)
	}
}

func TestDecode(t *testing.T) {
	input := "Gb trg gb gur bgure fvqr!"
	expected := "To get to the other side!"
	output := Decode(input)

	if output != expected {
		t.Errorf("Expected: %q Got: %q Using %q", expected, output, input)
	}
}

package main

import "testing"

func Test_createSectionFromString(t *testing.T) {

	original := "Gerald"

	section := createSectionFromString(original)

	if section.source != original {
		t.Errorf("Source does not match the original string. got: %s, want: %s", section.source, original)

	}

	if section.start != 0 {
		t.Errorf("Start is not correct. got: %s, want: %s", section.end, 0)
	}

	if section.end != len(original) {
		t.Errorf("End is not correct. got: %s, want: %s", section.end, len(original))
	}

	if section.toString() != original {
		t.Errorf("toString() is not correct. got: %s, want: %s", section.toString(), original)
	}

}
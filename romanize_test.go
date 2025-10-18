package cantonese

import (
	"reflect"
	"testing"
)

func TestRomanize(t *testing.T) {
	results := Romanize("香港😀", SchemeRoman)
	if len(results) != 3 {
		t.Fatalf("Romanize length = %d, want 3", len(results))
	}

	first := results[0]
	if first.Character != "香" || !first.Exists {
		t.Fatalf("unexpected first result: %+v", first)
	}

	wantPron := []string{"heung", "hong"}
	if !reflect.DeepEqual(first.Pronunciations, wantPron) {
		t.Fatalf("pronunciations = %v, want %v", first.Pronunciations, wantPron)
	}

	unknown := results[2]
	if unknown.Exists {
		t.Fatalf("expected unknown character to have Exists=false, got %+v", unknown)
	}
	if len(unknown.Pronunciations) != 0 {
		t.Fatalf("expected unknown character to have no pronunciations, got %+v", unknown)
	}

	t.Logf("LSHK: %v", RomanizeFirst("香港", SchemeLSHK))
	t.Logf("YALE: %v", RomanizeFirst("香港", SchemeYale))
	t.Logf("ROMAN: %v", RomanizeFirst("香港", SchemeRoman))

}

func TestRomanizeFirst(t *testing.T) {
	got := RomanizeFirst("香港", SchemeLSHK)
	want := "hoeng1 gong2"
	if got != want {
		t.Fatalf("RomanizeFirst = %q, want %q", got, want)
	}

}

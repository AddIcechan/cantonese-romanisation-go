package cantonese

import (
	"reflect"
	"testing"
)

func TestLookup(t *testing.T) {
	entry, ok := Lookup('万')
	if !ok {
		t.Fatalf("Lookup('万') returned ok=false")
	}

	wantRoman := []string{"mak", "maan", "baak", "bark", "pak"}
	if !reflect.DeepEqual(entry.Roman, wantRoman) {
		t.Fatalf("Lookup('万').Roman = %v, want %v", entry.Roman, wantRoman)
	}

	wantLSHK := []string{"maak6", "maan6", "mak6"}
	if !reflect.DeepEqual(entry.LSHK, wantLSHK) {
		t.Fatalf("Lookup('万').LSHK = %v, want %v", entry.LSHK, wantLSHK)
	}
}

func TestRomanizeRune(t *testing.T) {
	got, ok := RomanizeRune('与', SchemeLSHK)
	if !ok {
		t.Fatalf("RomanizeRune('与', SchemeLSHK) returned ok=false")
	}

	want := []string{"jyu4", "jyu5", "jyu6"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("RomanizeRune('与', SchemeLSHK) = %v, want %v", got, want)
	}
}

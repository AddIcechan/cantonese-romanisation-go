package cantonese

import "strings"

// Romanization captures the pronunciations available for an individual rune.
type Romanization struct {
	Rune           rune
	Character      string
	Pronunciations []string
	Exists         bool
}

// Romanize returns the pronunciations for each rune in the provided text using the
// chosen scheme. Unknown characters are still included in the result with
// Exists equal to false.
func Romanize(text string, scheme Scheme) []Romanization {
	ensureDictionary()
	results := make([]Romanization, 0, len(text))

	for _, r := range text {
		entry, ok := entries[r]
		var pronunciations []string
		if ok {
			pronunciations = entry.Romanisations(scheme)
		}

		results = append(results, Romanization{
			Rune:           r,
			Character:      string(r),
			Pronunciations: pronunciations,
			Exists:         ok,
		})
	}

	return results
}

// RomanizeFirst returns a space-delimited string choosing the first pronunciation
// (if available) for every rune in the input string and falling back to the raw
// character when no pronunciation is known.
func RomanizeFirst(text string, scheme Scheme) string {
	results := Romanize(text, scheme)
	output := make([]string, 0, len(results))

	for _, res := range results {
		if len(res.Pronunciations) > 0 {
			output = append(output, res.Pronunciations[0])
		} else {
			output = append(output, res.Character)
		}
	}

	return strings.Join(output, " ")
}

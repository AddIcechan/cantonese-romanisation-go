package cantonese

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"
)

// Scheme represents one of the supported Cantonese romanisation systems
// provided by the embedded dictionary.
type Scheme string

const (
	// SchemeRoman is the raw romanisation column from the source data.
	SchemeRoman Scheme = "roman"
	// SchemeLSHK corresponds to the Linguistic Society of Hong Kong (Jyutping) romanisation.
	SchemeLSHK Scheme = "lshk"
	// SchemeYale corresponds to the Yale romanisation.
	SchemeYale Scheme = "yale"
)

// Valid reports whether the scheme matches one of the supported constants.
func (s Scheme) Valid() bool {
	switch s {
	case SchemeRoman, SchemeLSHK, SchemeYale:
		return true
	default:
		return false
	}
}

// Entry contains the known pronunciations for a single Han character.
type Entry struct {
	Character string
	Roman     []string
	LSHK      []string
	Yale      []string
}

// Romanisations returns the pronunciations for the chosen scheme.
func (e Entry) Romanisations(s Scheme) []string {
	switch s {
	case SchemeRoman:
		return e.Roman
	case SchemeLSHK:
		return e.LSHK
	case SchemeYale:
		return e.Yale
	default:
		return nil
	}
}

//go:embed dictionary/source.tsv
var dictionaryTSV string

var (
	loadOnce sync.Once
	loadErr  error
	entries  map[rune]Entry
)

func ensureDictionary() {
	loadOnce.Do(func() {
		entries = make(map[rune]Entry, 14000)
		scanner := bufio.NewScanner(strings.NewReader(dictionaryTSV))

		line := 0
		for scanner.Scan() {
			raw := scanner.Text()
			line++

			if line == 1 {
				// Header row.
				continue
			}

			if strings.TrimSpace(raw) == "" {
				continue
			}

			fields := strings.Split(raw, "\t")
			if len(fields) != 4 {
				loadErr = fmt.Errorf("dictionary: line %d: expected 4 columns, got %d", line, len(fields))
				return
			}

			character := fields[0]
			if character == "" {
				loadErr = fmt.Errorf("dictionary: line %d: empty character", line)
				return
			}

			r, width := utf8.DecodeRuneInString(character)
			if r == utf8.RuneError && width == 0 {
				loadErr = fmt.Errorf("dictionary: line %d: invalid rune", line)
				return
			}

			entry := Entry{
				Character: character,
				Roman:     parsePronunciations(fields[1]),
				LSHK:      parsePronunciations(fields[2]),
				Yale:      parsePronunciations(fields[3]),
			}

			entries[r] = entry
		}

		if err := scanner.Err(); err != nil {
			loadErr = fmt.Errorf("dictionary: read error: %w", err)
		}
	})

	if loadErr != nil {
		panic(loadErr)
	}
}

func parsePronunciations(column string) []string {
	column = strings.TrimSpace(column)
	if column == "" {
		return nil
	}
	return strings.Fields(column)
}

// Lookup returns the dictionary entry for the provided rune, if available.
func Lookup(r rune) (Entry, bool) {
	ensureDictionary()
	entry, ok := entries[r]
	return entry, ok
}

// RomanizeRune returns the pronunciations for a single rune and reports
// whether an entry exists for that rune.
func RomanizeRune(r rune, scheme Scheme) ([]string, bool) {
	ensureDictionary()
	entry, ok := entries[r]
	if !ok {
		return nil, false
	}
	return entry.Romanisations(scheme), true
}

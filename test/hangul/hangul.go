// Package hangul provides basic functions for Hangul processing.package hangul
package hangul

var (
	start = rune(44032)
	end   = rune(55204)
)

// HasConsonantSuffix returns true if s has Hangul consonant jamo at the end.
func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

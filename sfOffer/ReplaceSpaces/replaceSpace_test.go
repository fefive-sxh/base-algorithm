package ReplaceSpaces

import "testing"

type strReplace struct {
	old    string
	expect string
}

func createStr(t *testing.T, s *strReplace) {
	t.Helper()
	if result := ReplaceSpaces(s.old); result != s.expect {
		t.Fatalf("%v expect: %v but got %v", s.old, s.expect, result)
	}
}

func TestReplaceSpaces(t *testing.T) {
	createStr(t, &strReplace{old: "We are happy.", expect: "We%20are%20happy."})
	createStr(t, &strReplace{old: " .", expect: "%20."})
	createStr(t, &strReplace{old: "  ", expect: "%20%20"})
}

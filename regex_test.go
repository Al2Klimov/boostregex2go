package boostregex2go

import (
	"reflect"
	"testing"
)

func TestRegex_MatchesSomewhere(t *testing.T) {
	assertRegex_MatchesSomewhere(t, "(", BadPattern{}, "", nil, false)
	assertRegex_MatchesSomewhere(t, "ab*+c", nil, "abc", nil, true)
	assertRegex_MatchesSomewhere(t, "ab*+c", nil, "adc", nil, false)
}

func assertRegex_MatchesSomewhere(t *testing.T, pattern string, compileError error, subject string, matchError error, matches bool) {
	t.Helper()

	rgx, errRgx := NewRegex([]byte(pattern))
	if rgx != nil {
		defer rgx.Close()
	}

	if !reflect.DeepEqual(errRgx, compileError) {
		t.Errorf("_, err := NewRegex([]byte(%#v)); !reflect.DeepEqual(err, %#v)", pattern, compileError)
	} else if compileError == nil {
		if match, errMatch := rgx.MatchesSomewhere([]byte(subject)); !reflect.DeepEqual(errMatch, matchError) {
			t.Errorf(
				"rgx, _ := NewRegex([]byte(%#v)); _, err := rgx.MatchesSomewhere([]byte(%#v)); !reflect.DeepEqual(err, %#v)",
				pattern, subject, matchError,
			)
		} else if match != matches {
			t.Errorf(
				"rgx, _ := NewRegex([]byte(%#v)); match, _ := rgx.MatchesSomewhere([]byte(%#v)); match != %#v",
				pattern, subject, matches,
			)
		}
	}
}

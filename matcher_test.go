package btcvanity

import "testing"

func TestMatcher(t *testing.T) {

	testCases := map[string]struct {
		Pattern  string
		Expected bool
	}{
		"1HvA9A8mVVJMhdLduC4NZpV4JQimNR3n4c": {"h", true},
		"1HvA9A8mVVJMhdLduC4NZpV4JQimNR3n4d": {"hva", true},
		"1HvA9A8mVVJMhdLduC4NZpV4JQimNR3n4e": {"Z", false},
		"HvA9A8mVVJMhdLduC4NZpV4JQimNR3n4c":  {"H", false},
	}

	for key, val := range testCases {
		got := isMatch(val.Pattern, key)
		if got != val.Expected {
			t.Errorf("error in test case %v. Expected %v got %v\n", key, val.Expected, got)
		}
	}
}

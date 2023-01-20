package fuzzy

import "testing"

func TestFuzzy(t *testing.T) {
    cases := []struct{
        matcher string
        input string
        expected bool 
    }{
        {"a", "ab", true},
        {"aa", "a", false},
        {"c", "ab", false},
        {"a","ba", true},
        {"ab", "abc", true},
        {"ac", "abc", true},
        {"bb", "abc", false},
        {"zo", "zero", true},
        {"za", "zero", false},
        {"zz", "zero", false},
    }

    for _, c := range cases {
        res := Matches(c.matcher, c.input);

        if res != c.expected {
            t.Fatalf("Expected %v, got %v for %s -> %s\n", c.expected, res, c.matcher, c.input)
        }
    }
}

package fuzzy

import "strings"

func Matches(matcher, input string) bool {
    for _, c := range matcher {
        if !strings.ContainsRune(input, c) {
            return false
        }

        idx := strings.IndexRune(input, c)
        a := input[:idx]
        b := input[idx+1:]
        input = a + b
    }
    return true
}

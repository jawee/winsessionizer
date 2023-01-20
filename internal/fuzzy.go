package fuzzy

import "strings"

func Matches(a, b string) bool {
    return strings.Contains(a, b)
}

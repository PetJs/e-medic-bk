// Package postgres provides PostgreSQL database implementations.
package postgres

import "strings"

// escapeLike escapes ILIKE metacharacters in a user-supplied search term so
// it's matched literally rather than as a wildcard pattern (e.g. searching
// for a literal "50%" shouldn't become "matches anything containing 50").
// Callers wrap the result in "%"..."%" and pass it as a bound parameter —
// this is a correctness fix, not an injection concern (the value is always
// parameterized either way).
func escapeLike(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)
	return r.Replace(s)
}

package envsubst

import (
	"github.com/hguerra/envsubst/v2/provider"
)

// Eval replaces ${var} in the string based on the mapping function.
func Eval(s string, mapping func(string) string) (string, error) {
	t, err := Parse(s)
	if err != nil {
		return s, err
	}
	return t.Execute(mapping)
}

// EvalEnv replaces ${var} in the string according to the values of the
// current environment variables. References to undefined variables are
// replaced by the empty string.
func EvalEnv(s string, noEmpty bool) (string, error) {
	mapping := provider.Get
	if noEmpty {
		mapping = provider.GetRequired
	}
	return Eval(s, mapping)
}

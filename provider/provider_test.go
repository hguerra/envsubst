package provider

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	os.Setenv("TESTGET", "myvar")

	var expressions = []struct {
		input  string
		output string
	}{
		{
			input:  "TESTGET",
			output: "myvar",
		},
		{
			input:  "gcp:secretmanager:projects/xxx/secrets/mykey/versions/1",
			output: "",
		},
	}

	for _, expr := range expressions {
		t.Run(expr.input, func(t *testing.T) {
			t.Logf(expr.input)
			output := Get(expr.input)
			if output != expr.output {
				t.Errorf("Want %q expanded to %q, got %q",
					expr.input,
					expr.output,
					output)
			}
		})
	}
}

func Test_isSecretManager(t *testing.T) {
	var expressions = []struct {
		input  string
		output bool
	}{
		{
			input:  "abcdEFGH28ij",
			output: false,
		},
		{
			input:  "gcp:secretmanager:projects/xxx/secrets/mykey/versions/1",
			output: true,
		},
	}

	for _, expr := range expressions {
		t.Run(expr.input, func(t *testing.T) {
			t.Logf(expr.input)
			output := isSecretManager(expr.input)
			if output != expr.output {
				t.Errorf("Want %v expanded to %v, got %v",
					expr.input,
					expr.output,
					output)
			}
		})
	}
}

func Test_getKey(t *testing.T) {
	var expressions = []struct {
		input  string
		output string
	}{
		{
			input:  "abcdEFGH28ij",
			output: "abcdEFGH28ij",
		},
		{
			input:  "gcp:secretmanager:projects/xxx/secrets/mykey/versions/1",
			output: "projects/xxx/secrets/mykey/versions/1",
		},
	}

	for _, expr := range expressions {
		t.Run(expr.input, func(t *testing.T) {
			t.Logf(expr.input)
			output := getKey(expr.input)
			if output != expr.output {
				t.Errorf("Want %v expanded to %v, got %v",
					expr.input,
					expr.output,
					output)
			}
		})
	}
}

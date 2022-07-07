package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ppcamp/go-lib/env"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	assert := require.New(t)

	expectedInt := 3
	expectedStr := "some_a"
	expectedFloat := 1.31

	envVars := map[string]string{
		"SOME_INT":    fmt.Sprintf("%d", expectedInt),
		"SOME_FLOAT":  fmt.Sprintf("%f", expectedFloat),
		"SOME_STRING": expectedStr,
	}

	// update this environ
	for key, value := range envVars {
		if err := os.Setenv(key, value); err != nil {
			t.Fail()
		}
	}

	var resultInt int
	var resultFloat float64
	var resultStr string

	// try to get those parses
	flags := []env.Flag{
		&env.BaseFlag[int]{Value: &resultInt, EnvName: "SOME_INT"},
		&env.BaseFlag[float64]{Value: &resultFloat, EnvName: "SOME_FLOAT"},
		&env.BaseFlag[string]{Value: &resultStr, EnvName: "SOME_STRING"},
	}

	// check if occurred some error during parse
	assert.NotPanics(func() {
		env.Parse(flags)
	})

	// check if parsed successfully
	results := []any{resultInt, resultFloat, resultStr}
	expected := []any{expectedInt, expectedFloat, expectedStr}
	assert.ElementsMatch(results, expected)
}

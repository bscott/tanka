package jsonnet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTransitiveImports checks that TransitiveImports is able to report all
// recursive imports of a file
func TestTransitiveImports(t *testing.T) {
	imports, err := TransitiveImports("testdata/environments/test/main.jsonnet")
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{
		"testdata/environments/test/trees.jsonnet",

		"testdata/environments/test/trees/apple.jsonnet",
		"testdata/environments/test/trees/cherry.jsonnet",
		"testdata/environments/test/trees/peach.jsonnet",

		"testdata/environments/test/trees/generic.libsonnet",
	}, imports)
}

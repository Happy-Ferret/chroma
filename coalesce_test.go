package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCoalesce(t *testing.T) {
	lexer := Coalesce(MustNewLexer(nil, Rules{
		"root": []Rule{
			Rule{`[!@#$%^&*()]`, Punctuation, nil},
		},
	}))
	actual, err := Tokenise(lexer, nil, "!@#$")
	require.NoError(t, err)
	expected := []*Token{{Punctuation, "!@#$"}}
	require.Equal(t, expected, actual)
}

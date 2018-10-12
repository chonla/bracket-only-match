package bracket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var MethodUnderTest = Match2

func TestEmpty(t *testing.T) {
	result := MethodUnderTest("")
	assert.True(t, result)
}

func TestCorrect(t *testing.T) {
	open := []string{"[]", "()", "{}", "[][]", "[](){}", "[({})][{}()]"}
	for _, o := range open {
		result := MethodUnderTest(o)
		assert.True(t, result)
	}
}

func TestOpenOnly(t *testing.T) {
	open := []string{"[", "(", "{"}
	for _, o := range open {
		result := MethodUnderTest(o)
		assert.False(t, result)
	}
}

func TestCloseOnly(t *testing.T) {
	close := []string{"]", ")", "}"}
	for _, o := range close {
		result := MethodUnderTest(o)
		assert.False(t, result)
	}
}

func TestCloseThenOpen(t *testing.T) {
	closeThenOpen := []string{"][", ")(", "}{"}
	for _, o := range closeThenOpen {
		result := MethodUnderTest(o)
		assert.False(t, result)
	}
}

func TestMixed(t *testing.T) {
	mixed := []string{"[]]", "{{(}})", "{[}(])"}
	for _, o := range mixed {
		result := MethodUnderTest(o)
		assert.False(t, result)
	}
}

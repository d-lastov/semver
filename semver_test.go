package semver_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.True(t, IsValid("1.33.4"), "expected 1.33.4 be valid")
	assert.True(t, IsValid("0.33.4"), "expected 0.33.4 be valid")
	assert.True(t, IsValid("0.0.4"), "expected 0.0.4 be valid")
	assert.True(t, IsValid("0.0.0"), "expected 0.0.0 be valid")
	assert.False(t, IsValid("-1.33.4"), "expected -1.33.4 be invalid")
}

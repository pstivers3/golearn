// https://www.youtube.com/watch?v=Kf5LLi7Ti9A&t=12s

// To run, go test
// Assert should fail with passord of at least 5 characters.
// However it passes.

import(
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestPassword_Validate(t *testing.T) {
	p := Password("321123")
	err := p.Validate()
	assert.Error(t, err)
}
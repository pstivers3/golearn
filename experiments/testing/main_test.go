import(
	"testing"
	"github.com/stretchr/testify/assert"
)

// assert should fail, however it passes

func TestPassword_Validate(t *testing.T) {
	p := Password("321123")
	err := p.Validate()
	assert.Error(t, err)
}
package rand_test

import (
	"math"
	"testing"

	"github.com/romv7/blogs/internal/utils/rand"
)

func Test_shouldGenerateARandomNumberFrom0To32(t *testing.T) {
	n, _ := rand.Rand()

	if n > uint64(math.Pow(2, 64)) || n < 0 {
		t.Errorf("error: Rand(n) did not follow the range [0, 2^32)")
	}
}

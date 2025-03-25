package float16_test

import (
	"math/rand"
	"testing"

	. "github.com/samborkent/float16"
	f16 "github.com/x448/float16"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	a1 := rand.Float32()
	a2 := rand.Float32()

	b1 := Float16{Float16: f16.Fromfloat32(a1)}
	b2 := Float16{Float16: f16.Fromfloat32(a2)}

	t.Logf("a1: %g, b1: %s", a1, b1.String())
	t.Logf("a2: %g, b2: %s", a2, b2.String())

	t.Logf("a sum: %g, b sum: %s", a1+a2, b1.Add(b2).String())
}

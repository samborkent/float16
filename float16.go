package float16

import (
	f16 "github.com/x448/float16"
)

type Float16 struct {
	f16.Float16
}

// Frombits returns the float16 number corresponding to the IEEE 754 binary16
// representation u16, with the sign bit of u16 and the result in the same bit
// position. Frombits(Bits(x)) == x.
func Frombits(u16 uint16) Float16 {
	return Float16{Float16: f16.Frombits(u16)}
}

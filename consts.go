package float16

var (
	Zero    = Float16{Float16: 0b0_00000_0000000000}
	NegZero = Float16{Float16: 0b1_00000_0000000000}
	One     = Float16{Float16: 0b0_01111_0000000000}
	NegOne  = Float16{Float16: 0b1_01111_0000000000}
)

const (
	bitLen      = 16
	sgnLen      = 1
	expLen      = 5
	mantissaLen = 10

	bias        = 0b01111
	maxExp      = 0b0_11110
	maxMantissa = (1 << mantissaLen) - 1

	maxFloat16 = 0b0_11110_1111111111
	minFloat16 = 0b1_11110_1111111111
)

package float16

func (f Float16) Sign() uint16 {
	return f.Bits() >> (bitLen - sgnLen)
}

func (f Float16) WithSign(sign uint16) Float16 {
	f = Frombits((f.Bits() | 0b0_11111_1111111111) | ((sign & 0b1) << (expLen + mantissaLen)))
	return f
}

func (f Float16) Exp() uint16 {
	return ((f.Bits() >> mantissaLen) & 0b0_11111) - bias
}

func (f Float16) WithExp(exp uint16) Float16 {
	f = Frombits((f.Bits() | 0b1_00000_1111111111) | (((exp + bias) & 0b0_11111) << mantissaLen))
	return f
}

func (f Float16) Mantissa() uint16 {
	return (f.Bits() << (sgnLen + expLen)) >> (sgnLen + expLen)
}

func (f Float16) WithMantissa(mantissa uint16) Float16 {
	f = Frombits((f.Bits() | 0b1_11111_0000000000) | (mantissa & 0b0_00000_1111111111))
	return f
}

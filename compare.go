package float16

func (f Float16) Equal(x Float16) bool {
	return f.Bits() == x.Bits()
}

// func compare(a, b uint16)

func (f Float16) Less(x Float16) bool {
	// Compare sign bits
	if f.Sign() < x.Sign() {
		// f is positive, x is negative
		return false
	} else if f.Sign() > x.Sign() {
		// f is negative, x is positive
		return true
	}

	// Sign bits are equal
	// Compare (biased) exponents
	if f.Exp() < x.Exp() {
		return true
	} else if f.Exp() > x.Exp() {
		return false
	}

	// Exponents are equal
	// Compare mantissas
	return f.Mantissa() < x.Mantissa()
}

func (f Float16) LessEqual(x Float16) bool {
	// Compare sign bits
	if f.Sign() < x.Sign() {
		// f is positive, x is negative
		return false
	} else if f.Sign() > x.Sign() {
		// f is negative, x is positive
		return true
	}

	// Sign bits are equal
	// Compare (biased) exponents
	if f.Exp() < x.Exp() {
		return true
	} else if f.Exp() > x.Exp() {
		return false
	}

	// Exponents are equal
	// Compare mantissas
	return f.Mantissa() <= x.Mantissa()
}

func (f Float16) Greater(x Float16) bool {
	// Compare sign bits
	if f.Sign() > x.Sign() {
		// f is negative, x is positive
		return false
	} else if f.Sign() < x.Sign() {
		// f is positive, x is negative
		return true
	}

	// Sign bits are equal
	// Compare (biased) exponents
	if f.Exp() > x.Exp() {
		return true
	} else if f.Exp() < x.Exp() {
		return false
	}

	// Exponents are equal
	// Compare mantissas
	return f.Mantissa() > x.Mantissa()
}

func (f Float16) GreaterEqual(x Float16) bool {
	// Compare sign bits
	if f.Sign() > x.Sign() {
		// f is negative, x is positive
		return false
	} else if f.Sign() < x.Sign() {
		// f is positive, x is negative
		return true
	}

	// Sign bits are equal
	// Compare (biased) exponents
	if f.Exp() > x.Exp() {
		return true
	} else if f.Exp() < x.Exp() {
		return false
	}

	// Exponents are equal
	// Compare mantissas
	return f.Mantissa() >= x.Mantissa()
}

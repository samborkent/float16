package float16

func (f Float16) Add(x Float16) Float16 {
	if f.Exp() > x.Exp() {
		x = x.WithExp(f.Exp()).WithMantissa(x.Mantissa() >> (f.Exp() - x.Exp()))
	} else if f.Exp() < x.Exp() {
		f = f.WithExp(x.Exp()).WithMantissa(f.Mantissa() >> (x.Exp() - f.Exp()))
	}

	if f.Signbit() == x.Signbit() {
		// Both numbers have the same sign
		newMantissa := f.Mantissa() + x.Mantissa()
		if newMantissa > maxMantissa {
			newExp := ((f.Exp() + 1) + bias) << mantissaLen
			if newExp > maxExp {
				return Float16{Float16: maxFloat16}
			}

			return f.WithExp(newExp).WithMantissa(newMantissa)
		} else {
			return f.WithMantissa(newMantissa)
		}
	} else if !f.Signbit() && x.Signbit() {
		// f is positive, and x is negative
		if f.Mantissa() == x.Mantissa() {
			return Float16{Float16: 0}
		} else if f.Mantissa() < x.Mantissa() {
			// x mantissa is larger than f
			return f.WithSign(1).WithMantissa(x.Mantissa() - f.Mantissa())
		} else {
			// f mantissa is larger than x
			return f.WithMantissa(f.Mantissa() - x.Mantissa())
		}
	} else {
		// f is negative, and x is positive
		if f.Mantissa() == x.Mantissa() {
			return Float16{Float16: 0}
		} else if f.Mantissa() < x.Mantissa() {
			// x mantissa is larger than f
			return f.WithSign(0).WithMantissa(x.Mantissa() - f.Mantissa())
		} else {
			// f mantissa is larger than x
			return f.WithMantissa(f.Mantissa() - x.Mantissa())
		}
	}
}

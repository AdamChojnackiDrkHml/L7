package utilities

func CycleHamming() [16]uint8 {
	var h [16]uint8
	for i := uint8(0); i < 16; i++ {
		code := uint8(0)
		a := uint8((i >> 3) & 1)
		b := uint8((i >> 2) & 1)
		c := uint8((i >> 1) & 1)
		d := uint8(i & 1)
		code = uint8((a << 7) +
			(((a + b) % 2) << 6) +
			(((b + c) % 2) << 5) +
			(((a + c + d) % 2) << 4) +
			(((b + d) % 2) << 3) +
			(c << 2) +
			(d << 1))
		ones_counter := uint8(0)
		for j := 7; j > 0; j-- {
			if ((code >> j) & 1) == 1 {
				ones_counter++
			}
		}
		code += (ones_counter % 2)
		h[i] = code
	}
	return h
}

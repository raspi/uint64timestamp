package internal

// GetDigit fetches Nth digit
// indexing goes from right to left ( <-- )
func GetDigit(number uint64, n uint8) uint8 {
	return uint8(number / p10(n) % 10)
}

// p10 is Pow10 (power of 10)
func p10(n uint8) (o uint64) {
	o = 1

	for i := uint8(0); i < n; i++ {
		o *= 10
	}

	return o
}

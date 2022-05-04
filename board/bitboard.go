package board

// SetBit sets the bit at the given square to 1.
// The function returns a new bitboard.
func SetBit(bitboard uint64, square uint8) uint64 {
	return bitboard | (uint64(1) << square)
}

// GetBit returns the bit set at given square.
func GetBit(bitboard uint64, square uint8) uint8 {
	if bitboard&(uint64(1)<<square) == 0 {
		return 0
	}
	return 1
}

// ResetBit resets the bit at the given square to 1.
// The function returns a new bitboard.
func ResetBit(bitboard uint64, square uint8) uint64 {
	return bitboard & ^(uint64(1) << square)
}

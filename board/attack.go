package board

// MaskPawnAttacks returns a mask bitboard for pawn attack squares.
func MaskPawnAttacks(square uint8, side uint8) uint64 {
	bitboard := SetBit(uint64(0), square)

	// Checking for side because pawns can only move forward based on their color.
	// Also ignoring illegal captures moves that end on one rank up.
	if side == White {
		return ((bitboard >> 7) & ^FileA) | ((bitboard >> 9) & ^FileH)
	} else {
		return ((bitboard << 7) & ^FileH) | ((bitboard << 9) & ^FileA)
	}
}

// MaskKnightAttacks returns a mask bitboard for knight attack squares.
func MaskKnightAttacks(square uint8) uint64 {
	bitboard := SetBit(uint64(0), square)

	// computing the mask bitboard while ignoring overflow into illegal squares.
	// for example, when calculating bitboard >> 6, it simulates the move one up two right.
	// if the move ends on the AB files, it means we shifted one rank up and need to ignore the move.
	return ((bitboard >> 17) & ^FileH) |
		((bitboard >> 15) & ^FileA) |
		((bitboard >> 10) & ^(FileG | FileH)) |
		((bitboard >> 6) & ^(FileA | FileB)) |
		((bitboard << 17) & ^FileA) |
		((bitboard << 15) & ^FileH) |
		((bitboard << 10) & ^(FileA | FileB)) |
		((bitboard << 6) & ^(FileG | FileH))
}

// MaskKingAttacks returns a mask bitboard for king attack squares.
func MaskKingAttacks(square uint8) uint64 {
	bitboard := SetBit(uint64(0), square)

	// computing the mask bitboard while ignoring overflow into illegal squares.
	// for example, when calculating bitboard >> 1, it simulates the move one left.
	// if the move ends on the H File, that means we shifted one rank up and need to ignore the move.
	return ((bitboard >> 9) & ^FileH) |
		(bitboard >> 8) |
		((bitboard >> 7) & ^FileA) |
		((bitboard >> 1) & ^FileH) |
		((bitboard << 9) & ^FileH) |
		(bitboard << 8) |
		((bitboard << 7) & ^FileH) |
		((bitboard << 1) & ^FileA)
}

// MaskBishopAttacks returns a mask bitboard for bishop attack squares.
// Squared on the board's edges are not included in the mask bitboard.
func MaskBishopAttacks(square uint8) uint64 {
	//computing the rank and file of the square.
	tr := square / 8
	tf := square % 8

	// going from the square down on the main diagonal.
	// there is no underflow in this case.
	main_down := uint64(0)
	for rank, file := tr+1, tf+1; rank <= 6 && file <= 6; rank, file = rank+1, file+1 {
		main_down = SetBit(main_down, rank*8+file)
	}

	// going from the square up on the anti-diagonal.
	// there is no underflow in this case.
	anti_up := uint64(0)
	for rank, file := tr-1, tf+1; rank >= 1 && file <= 6; rank, file = rank-1, file+1 {
		anti_up = SetBit(anti_up, rank*8+file)
	}

	// going from the square down on the anti-diagonal.
	// checking if target file is not zero to prevent underflow.
	anti_down := uint64(0)
	if tf != 0 {
		for rank, file := tr+1, tf-1; rank <= 6 && file >= 1; rank, file = rank+1, file-1 {
			anti_down = SetBit(anti_down, rank*8+file)
		}
	}

	// going from the square up on the main diagonal.
	// checking if target file and target rank are not zero to prevent underflow.
	main_up := uint64(0)
	if tr != 0 && tf != 0 {
		for rank, file := tr-1, tf-1; rank >= 1 && file >= 1; rank, file = rank-1, file-1 {
			main_up = SetBit(main_up, rank*8+file)
		}
	}

	return anti_down | anti_up | main_down | main_up
}

// MaskRookAttacks returns a mask bitboard for rook attack sqaures.
// Sqaures on the board's edges are not included in the mask bitboard.
func MaskRookAttacks(square uint8) uint64 {
	// computing rank and file from sqaure,
	tr := square / 8
	tf := square % 8

	// calculating the mask bitboard for downward moves.
	down := uint64(0)
	for rank := tr + 1; rank < 7; rank++ {
		down = SetBit(down, rank*8+tf)
	}

	// calculating the mask bitboard for upward moves.
	up := uint64(0)
	for rank := tr - 1; rank > 0; rank-- {
		up = SetBit(up, rank*8+tf)
	}

	// calculating the mask bitboard for rightward moves.
	right := uint64(0)
	for file := tf + 1; file < 7; file++ {
		right = SetBit(right, tr*8+file)
	}

	// calculating the mask bitboard for leftward moves.
	left := uint64(0)
	for file := tf - 1; file > 0; file-- {
		left = SetBit(left, tr*8+file)
	}

	// ORing all the calculated bitboards to get the final bitboard.
	return down | up | right | left
}

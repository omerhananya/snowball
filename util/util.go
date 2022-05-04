package util

import (
	"fmt"

	"github.com/omerhananya/snowball/board"
)

// PrintBitboard prints bitboard
func PrintBitboard(bitboard uint64) {
	fmt.Printf("\n")
	for rank := 0; rank < 8; rank++ {

		// print ranks
		fmt.Printf("%v    ", 8-rank)

		for file := 0; file < 8; file++ {

			// calculate square index
			square := uint8(rank*8 + file)

			// get bit and print
			bit := board.GetBit(bitboard, square)
			fmt.Printf("%v ", bit)
		}

		fmt.Printf("\n")
	}
	fmt.Printf("\n     a b c d e f g h\n\n")
}

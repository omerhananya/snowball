package io

import (
	"fmt"

	"github.com/omerhananya/snowball/board"
)

var bit_to_char = [...]string{".", "1"}

// PrintBitboard prints bitboard
// The function recieved a bitboard of type uint64 and displays its state.
// 1s will be represented as 1s, while 0s will be represented as "." (dot)
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
			fmt.Printf("%v ", bit_to_char[bit])
		}
		fmt.Printf("\n")
	}
	// print column indexes
	fmt.Printf("\n     a b c d e f g h\n\n")
}

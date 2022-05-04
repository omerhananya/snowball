package board

// board positions
const (
	A8 uint8 = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

// sides to move (colors)
const (
	White uint8 = iota
	Black
)

// files bitboards
const (
	FileA uint64 = 72340172838076673
	FileB uint64 = 144680345676153346
	FileC uint64 = 289360691352306692
	FileD uint64 = 578721382704613384
	FileE uint64 = 1157442765409226768
	FileF uint64 = 2314885530818453536
	FileG uint64 = 4629771061636907072
	FileH uint64 = 9259542123273814144
)

// ranks bitboards
const (
	Rank1 uint64 = 18374686479671623680
	Rank2 uint64 = 71776119061217280
	Rank3 uint64 = 280375465082880
	Rank4 uint64 = 1095216660480
	Rank5 uint64 = 4278190080
	Rank6 uint64 = 16711680
	Rank7 uint64 = 65280
	Rank8 uint64 = 255
)

package chessboard

type File = []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard = []File{
	"A", File,
	
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}

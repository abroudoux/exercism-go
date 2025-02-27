package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	var count int

	for _, v := range cb[file] {
		if v {
			count++
		}
	}

	return count
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	var count int
	for _, v := range cb {
		if v[rank-1] {
			count++
		}
	}

	return count
}

func CountAll(cb Chessboard) int {
	var count int
	for _, v := range cb {
		count += len(v)
	}

	return count
}

func CountOccupied(cb Chessboard) int {
	var count int
	for _, v := range cb {
		for _, occupied := range v {
			if occupied {
				count++
			}
		}
	}

	return count
}

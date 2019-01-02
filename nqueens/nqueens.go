// Package nqueens provides solutions to the N-queens puzzle.
package nqueens

//func main() {
//	nPtr := flag.Int("n", 8, "Dimension of the board to give an n*n board. Defaults to 8.")
//	flag.Parse()
//
//	solutions := GetAllSolutions(*nPtr)
//	solutions = ReduceToFundamentalSolutions(solutions)
//	for i, b := range solutions {
//		println(i+1)
//		fmt.Print(b)
//	}
//}

// GetAllSolutions finds all of the distinct solutions for an NxN board.
func GetAllSolutions(start Board) []Board {
	return start.getAllSolutions()
}

func (b Board) getAllSolutions() []Board {
	// Find the first unoccupied row
	var y int
	for y = 0; y < b.N; y++ {
		if !b.unsafeRows[y] {
			break
		}
	}

	var allSolutions []Board

	// Base case for recursion
	if y == b.N {
		allSolutions = append(allSolutions, b)
		return allSolutions
	}

	for x := 0; x < b.N; x++ {
		if b.IsSafe(x, y) {
			b1 := b.SetQueen(x, y)

			// Recurse
			solutions := b1.getAllSolutions()
			allSolutions = append(allSolutions, solutions...)
		}
	}
	return allSolutions
}

// ReduceToFundamentalSolutions returns a subset of the given solutions, where
// all variants that differ by rotation or reflection are counted as one solution.
func ReduceToFundamentalSolutions(solutions []Board) []Board {
	var reducedSolutions []Board
	for len(solutions) > 0 {
		head := solutions[0]
		tail := solutions[1:]
		
		reducedSolutions = append(reducedSolutions, head)

		var reduced []Board
		for _, b := range tail {
			if !head.IsEquivalent(b) {
				reduced = append(reduced, b)
			}
		}
		solutions = reduced
	}
	return reducedSolutions
}

// IsEqual determines whether two given boards have the same placement of queens.
func (b Board) IsEqual(b1 Board) bool {
	if len(b.Queens) != len(b1.Queens) {
		return false
	}

	for i, q1 := range b.Queens {
		if q1 != b1.Queens[i] {
			return false
		}
	}
	return true
}

// IsEquivalent determines whether the placement of queens on one board is a variant of another.
// A variant can be the same or differs by rotating or mirroring the board. Such variants are
// counted as one for a fundamental solution.
func (b Board) IsEquivalent(b1 Board) bool {
	// Rotate
	for i := 0; i < 4; i++ {
		b1 = b1.RotateClockwise()
		if b.IsEqual(b1) {
			return true
		}
	}

	// Mirror & rotate
	b3 := b1.Mirror()
	for i := 0; i < 4; i++ {
		if b.IsEqual(b3) {
			return true
		}
		b3 = b3.RotateClockwise()
	}

	return false
}

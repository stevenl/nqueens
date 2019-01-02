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

// GetAllSolutions finds all of the distinct solutions for the given board.
// If the board is empty then it will return all possible solutions for that
// board size.
// If the board is preset with starting queens then it will return  only the
// solutions that include those positions being occupied.
// If the given board is preset with queens that already threaten each
// other then no solutions will be returned.
func GetAllSolutions(start Board) []Board {
	if !start.isValid() {
		return make([]Board, 0)
	}
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

	// Recursively find the columns that are possible for a solution
	for x := 0; x < b.N; x++ {
		if b.IsSafe(x, y) {
			b1 := b.SetQueen(x, y)
			solutions := b1.getAllSolutions()
			allSolutions = append(allSolutions, solutions...)
		}
	}
	return allSolutions
}

// ReduceToFundamentalSolutions returns a subset of the given solutions, where
// all variants that differ by rotation or reflection are counted as a single
// solution.
func ReduceToFundamentalSolutions(solutions []Board) []Board {
	var reducedSolutions []Board
	for len(solutions) > 0 {
		head := solutions[0]
		tail := solutions[1:]
		variants := head.variants()

		reducedSolutions = append(reducedSolutions, head)

		var reduced []Board
		for _, b := range tail {
			if !b.hasMatch(variants) {
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

// Variants returns a list of Boards that are variants of the original board.
// Variants can be the same or differ by rotating or mirroring the board.
// The original board configuration is included in the list of variants.
func (b Board) variants() []Board {
	var variants []Board

	// Rotate
	for i := 0; i < 4; i++ {
		variants = append(variants, b)
		b = b.RotateClockwise()
	}

	// Mirror & rotate
	b = b.Mirror()
	for i := 0; i < 4; i++ {
		variants = append(variants, b)
		b = b.RotateClockwise()
	}
	return variants
}

func (b Board) hasMatch(variants []Board) bool {
	for _, v := range variants {
		if b.IsEqual(v) {
			return true
		}
	}
	return false
}

func (b Board) isValid() bool {
	checker := NewBoard(b.N)
	for _, q := range b.Queens {
		if !checker.IsSafe(q.X, q.Y) {
			return false
		}
		checker.setQueen(q)
	}
	return true
}

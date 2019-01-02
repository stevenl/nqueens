package nqueens

import (
	"fmt"
	"testing"
)

func Example() {
	b := NewBoard(8)
	allSolutions := GetAllSolutions(b)
	for i, solution := range allSolutions {
		fmt.Println(i + 1)
		fmt.Println(solution)
	}

	fundamentalSolutions := ReduceToFundamentalSolutions(allSolutions)
	for i, solution := range fundamentalSolutions {
		fmt.Println(i + 1)
		fmt.Println(solution)
	}
}

func TestSolutions(t *testing.T) {
	var tests = []struct {
		N                      int
		NrAllSolutions         int
		NrFundamentalSolutions int
		NrSolutions            int
	}{
		{4, 2, 1, 1},
		{8, 92, 12, 16},
	}

	for _, test := range tests {
		solutions := GetAllSolutions(NewBoard(test.N))
		logSolutions(t, solutions)

		if len := len(solutions); len != test.NrAllSolutions {
			t.Errorf("GetAllSolutions(%d) = %d, want %d", test.N, len, test.NrAllSolutions)
		}

		solutions = ReduceToFundamentalSolutions(solutions)
		if len := len(solutions); len != test.NrFundamentalSolutions {
			t.Errorf("GetFundamentalSolutions(%d) = %d, want %d", test.N, len, test.NrFundamentalSolutions)
		}

		b := NewBoard(test.N).SetQueen(0, 2)
		solutions = GetAllSolutions(b)
		logSolutions(t, solutions)
		if len := len(solutions); len != test.NrSolutions {
			t.Errorf("GetSolutions(%d) = %d, want %d", test.N, len, test.NrSolutions)
		}
	}
}

func logSolutions(t *testing.T, solutions []Board) {
	for i, b := range solutions {
		t.Logf("Solution %d\n%s", i+1, b)
	}
}

func BenchmarkOverall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReduceToFundamentalSolutions(GetAllSolutions(NewBoard(8)))
	}
}

func BenchmarkGetAllSolutions(b *testing.B) {
	board := NewBoard(8)
	for i := 0; i < b.N; i++ {
		_ = GetAllSolutions(board)
	}
}

func BenchmarkReduceToFundamentalSolutions(b *testing.B) {
	allSolutions := GetAllSolutions(NewBoard(8))
	for i := 0; i < b.N; i++ {
		_ = ReduceToFundamentalSolutions(allSolutions)
	}
}

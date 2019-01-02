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

type Test struct {
	N                      int
	NrAllSolutions         int
	NrFundamentalSolutions int
}

func TestGetAllSolutions(t *testing.T) {
	var tests = []Test{
		{4, 2, 1},
		{8, 92, 12},
	}

	for _, test := range tests {
		solutions := testGetAllSolutions(t, test)
		testReduceToFundamentalSolutions(t, test, solutions)
	}
}

func testGetAllSolutions(t *testing.T, test Test) []Board {
	solutions := GetAllSolutions(NewBoard(test.N))
	t.Logf("%d solutions for a %dx%d board\n", len(solutions), test.N, test.N)
	logSolutions(t, solutions)

	if nrSolutions := len(solutions); nrSolutions != test.NrAllSolutions {
		t.Errorf("GetAllSolutions(%d) = %d, want %d", test.N, nrSolutions, test.NrAllSolutions)
	}
	return solutions
}

func testReduceToFundamentalSolutions(t *testing.T, test Test, allSolutions []Board) {
	solutions := ReduceToFundamentalSolutions(allSolutions)
	t.Logf("%d fundamental solutions for a %dx%d board\n", len(solutions), test.N, test.N)
	logSolutions(t, solutions)

	if nrSolutions := len(solutions); nrSolutions != test.NrFundamentalSolutions {
		t.Errorf("ReduceToFundamentalSolutions(%d) = %d, want %d", test.N, nrSolutions, test.NrFundamentalSolutions)
	}
}

func TestPresetBoards(t *testing.T) {
	tests := []struct {
		N             int
		InitPositions []position
		NrSolutions   int
	}{
		{4, []position{position{0, 2}}, 1},
		{4, []position{position{3, 3}}, 0},
		{8, []position{position{0, 2}}, 16},
		{8, []position{position{0, 2}, position{5, 4}}, 2},
		{8, []position{position{0, 2}, position{0, 0}}, 0},
	}

	for _, test := range tests {
		b := NewBoard(test.N)
		for _, p := range test.InitPositions {
			b = b.SetQueen(p.X, p.Y)
		}
		solutions := GetAllSolutions(b)
		t.Logf("%d solutions for a %dx%d board with initial positions set %s\n", len(solutions), test.N, test.N, b.Queens)
		logSolutions(t, solutions)

		if nrSolutions := len(solutions); nrSolutions != test.NrSolutions {
			t.Errorf("PresetBoards(%d%s) = %d, want %d", test.N, test.InitPositions, nrSolutions, test.NrSolutions)
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

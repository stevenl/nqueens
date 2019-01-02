package nqueens

import (
	"sort"
	"strconv"
	"strings"
)

// position consists of an X and Y coordinate that represents a position on the chess board.
type position struct {
	X, Y int
}

// byY type implements the Sorter interface to enable sorting of positions.
type byY []position

func (pp byY) Len() int           { return len(pp) }
func (pp byY) Swap(i, j int)      { pp[i], pp[j] = pp[j], pp[i] }
func (pp byY) Less(i, j int) bool { return pp[i].Y < pp[j].Y }

// Board represents a NxN chess board with a list of Queen positions on the board
type Board struct {
	N            int
	Queens       []position
	unsafeRows   map[int]bool
	unsafeCols   map[int]bool
	unsafeDiagUp map[int]bool
	unsafeDiagDn map[int]bool
}

// NewBoard creates a new, empty Board.
func NewBoard(n int) Board {
	b := Board{
		N:            n,
		unsafeRows:   make(map[int]bool),
		unsafeCols:   make(map[int]bool),
		unsafeDiagUp: make(map[int]bool),
		unsafeDiagDn: make(map[int]bool),
	}
	return b
}

// IsSafe determines whether the given position is safe from the queens currently on the board.
// It checks that there are no queens on the same row, column or diagonal.
func (b Board) IsSafe(x, y int) bool {
	isSafe := false
	switch {
	case b.unsafeRows[y]:
	case b.unsafeCols[x]:
	case b.unsafeDiagUp[y-x]:
	case b.unsafeDiagDn[y+x]:
	default:
		isSafe = true
	}
	return isSafe
}

func (b *Board) setQueen(q position) {
	b.Queens = append(b.Queens, q)
	b.unsafeRows[q.Y] = true
	b.unsafeCols[q.X] = true
	b.unsafeDiagUp[q.Y-q.X] = true
	b.unsafeDiagDn[q.Y+q.X] = true
}

// SetQueen copyies the board and sets a queen at the given position.
func (b Board) SetQueen(x, y int) Board {
	// Always keep the queens sorted
	needToSortQueens := false
	for y1 := range b.unsafeRows {
		if y1 > y {
			needToSortQueens = true
			break
		}
	}

	b = b.Copy()
	b.setQueen(position{x, y})

	if needToSortQueens {
		sort.Sort(byY(b.Queens))
	}
	return b
}

// Copy returns a copy of the board
func (b Board) Copy() Board {
	copy := NewBoard(b.N)
	for _, q := range b.Queens {
		copy.setQueen(q)
	}
	return copy
}

// Mirror returns a new board with the pieces reflected over the Y-axis
func (b Board) Mirror() Board {
	mirror := NewBoard(b.N)
	for _, q := range b.Queens {
		x := b.N - q.X - 1
		mirror.setQueen(position{x, q.Y})
	}
	return mirror
}

// RotateClockwise returns a new board with the pieces rotated clockwise by 90 degrees
func (b Board) RotateClockwise() Board {
	rotated := NewBoard(b.N)
	for _, q := range b.Queens {
		x := b.N - q.Y - 1
		rotated.setQueen(position{x, q.X})
	}
	sort.Sort(byY(rotated.Queens))
	return rotated
}

// String returns a string representation of the board.
// This method fulfils the requirement of the "fmt.Stringer" interface.
func (b Board) String() string {
	buildMatrix := func(b Board) [][]int {
		// Initialise empty matrix
		matrix := make([][]int, b.N)
		for y := range matrix {
			matrix[y] = make([]int, b.N)
		}

		// Set the queens in the matrix
		for _, q := range b.Queens {
			matrix[q.Y][q.X] = 1
		}
		return matrix
	}

	sb := strings.Builder{}

	// Top border
	sb.WriteString(" -")
	sb.WriteString(strings.Repeat("-", b.N*2))
	sb.WriteString("\n")

	matrix := buildMatrix(b)
	for _, row := range matrix {
		sb.WriteString("[ ")
		for _, col := range row {
			sb.WriteString(strconv.Itoa(col))
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}

	// Bottom border
	sb.WriteString(" -")
	sb.WriteString(strings.Repeat("-", b.N*2))
	sb.WriteString("\n")

	return sb.String()
}

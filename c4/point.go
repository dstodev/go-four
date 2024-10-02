package c4

type Point struct {
	Row    int
	Column int
}

func NewPoint(row, col int) Point {
	return Point{row, col}
}

func (p Point) Get() (row, col int) {
	return p.Row, p.Column
}

func (p Point) Step(direction Direction) Point {
	return p.Add(direction.Offset())
}

func (p Point) Add(other Point) Point {
	return Point{p.Row + other.Row, p.Column + other.Column}
}

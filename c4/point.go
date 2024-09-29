package c4

type Point struct {
	Row int
	Col int
}

func NewPoint(row, col int) Point {
	return Point{row, col}
}

func (p Point) Get() (row, col int) {
	return p.Row, p.Col
}

func (p Point) Step(direction Direction) Point {
	return p.Add(direction.Offset())
}

func (p Point) Add(other Point) Point {
	return Point{p.Row + other.Row, p.Col + other.Col}
}

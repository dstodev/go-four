package c4

type Point struct {
	Row int
	Col int
}

func (p Point) Get() (int, int) {
	return p.Row, p.Col
}

func (p Point) Step(direction Direction) Point {
	return Point{p.Row + direction.OffsetRow(), p.Col + direction.OffsetColumn()}
}

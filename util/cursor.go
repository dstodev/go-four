package util

type Cursor struct {
	row    int
	column int
}

func NewCursor() *Cursor {
	return &Cursor{}
}

func (c Cursor) Get() (row int, column int) {
	return c.Row(), c.Col()
}

func (c Cursor) Row() int {
	return c.row
}

func (c Cursor) Col() int {
	return c.column
}

func (c *Cursor) MoveUp() {
	c.row--
}

func (c *Cursor) MoveDown() {
	c.row++
}

func (c *Cursor) MoveLeft() {
	c.column--
}

func (c *Cursor) MoveRight() {
	c.column++
}

// Constrain the cursor between [min-row-inclusive, max-row-exclusive)
func (c *Cursor) ConstrainRow(min, max int) {
	constrain(&c.row, min, max)
}

// Constrain the cursor between [min-column-inclusive, max-column-exclusive)
func (c *Cursor) ConstrainCol(min, max int) {
	constrain(&c.column, min, max)
}

func constrain(value *int, min, max int) {
	if *value >= max {
		*value = max - 1
	}
	// Evaluate min after max to give min precedence
	if *value < min {
		*value = min
	}
}

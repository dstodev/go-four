package c4

type GameStatus int

func (status GameStatus) String() string {
	return [...]string{
		"Initial",
		"Running",
		"Completed",
		"Tied",
	}[status]
}

const (
	Initial GameStatus = iota
	Running
	Completed
	Draw
)

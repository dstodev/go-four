package c4

type GameStatus int

func (status GameStatus) String() string {
	return [...]string{
		"Initial",
		"Running",
		"Completed",
	}[status]
}

const (
	Initial GameStatus = iota
	Running
	Completed
)

package types

type Command int

const (
	Budget Command = iota
	Add
)

func (c Command) String() string {
	return [...]string{"budget", "add"}[c]
}

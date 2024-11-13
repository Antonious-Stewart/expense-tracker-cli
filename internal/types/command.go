package types

type Command int

const (
	Budget Command = iota
	Add
	List
)

func (c Command) String() string {
	return [...]string{"budget", "add", "list"}[c]
}

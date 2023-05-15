package graph

type Error struct{ s string }

func (err *Error) Error() string {
	return err.s
}

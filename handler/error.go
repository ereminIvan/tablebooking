package handler

type Error struct {
	Code  int
	Value string
}

func (e Error) Error() string {
	return e.Value
}

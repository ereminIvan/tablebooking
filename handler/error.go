package handler

type Error struct {
	Code  int
	Value string
}

func (e Error) Error() string {
	return e.Value
}

var (
	errInvalidGuestName     = Error{Value: "Invalid guest name"}
	errInvalidGuestLastName = Error{Value: "Invalid guest last name"}
	errInvalidEventTitle    = Error{Value: "Invalid event title"}
)

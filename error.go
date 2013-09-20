package osincli

// OAuth2 error base
type Error struct {
	Id          string
	Description string
	URI         string
	State       string
}

func (e *Error) Error() string {
	return e.Description
}

func NewError(id, description, uri, state string) *Error {
	return &Error{
		Id:          id,
		Description: description,
		URI:         uri,
		State:       state,
	}
}

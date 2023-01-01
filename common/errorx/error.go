package errorx

type MYError struct {
	s string
}

func New(s string) *MYError {
	return &MYError{s}
}

func (e *MYError) Error() string {
	return e.s
}

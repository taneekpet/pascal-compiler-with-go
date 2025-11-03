package models

const (
	Unrecognized = 1
)

type CustomError struct {
	ErrorIndex int
	Code       int
}

func (c CustomError) Error() string {
	return ""
}

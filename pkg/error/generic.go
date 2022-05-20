package error

import "encoding/json"

type GenericError struct {
	Message string
}

func New(err string) *GenericError {
	return &GenericError{
		Message: err,
	}
}

func (e *GenericError) ToJSON() string {
	j, _ := json.Marshal(e)
	return string(j)
}

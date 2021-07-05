package domain

import "fmt"

func HandleError(err error, src string) error {
	if err != nil {
		fmt.Println("[ keng demo][ Error] - ", err.Error())
	}
	return err
}

var _ error = &ServerError{}

type ServerError struct {
	Key          string
	DefaultError string
}

func NewError(key string, msg string) *ServerError {
	return &ServerError{
		Key:          key,
		DefaultError: msg,
	}
}

func (de *ServerError) Error() string {
	return de.DefaultError
}

func (de *ServerError) Set(msg string) {
	de.DefaultError = msg
}

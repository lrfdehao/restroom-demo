package context

import "fmt"

func HandleError(err error, src string) error {
	if err != nil {
		fmt.Println("[ toilet demo][ Error] - ", err.Error())
	}
	return err
}

var _ error = &ServerError{}

type ServerError struct {
	Key          int
	DefaultError string
}

func NewError(key int, msg string) *ServerError {
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

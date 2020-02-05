package router

import (
	"fmt"

	"github.com/suzuito/bookstore-go/entity"
)

var (
	NotFoundError         = fmt.Errorf("Not found")
	InvalidParameterError = fmt.Errorf("Invalid parameters")
)

type WrappedError struct {
	msg string
	err error
}

func NewWrappedError(err error, format string, a ...interface{}) *WrappedError {
	return &WrappedError{
		msg: fmt.Sprintf(format, a...),
		err: err,
	}
}

func (e *WrappedError) Unwrap() error { return e.err }
func (e *WrappedError) Error() string { return e.msg }

type Repository interface {
	GetBooks(book *[]*entity.Book) error
	GetBookByID(id string, book *entity.Book) error
}

package repository

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

func (e *WrappedError) UnWrap() error { return e.err }

type Repository interface {
	GetBooks(book *[]*entity.Book) error
	GetBookByID(id string, book *entity.Book) error
}

package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/bookstore-go/entity"
)

type responseError struct {
	Message string `json:"message"`
}

type responseBook struct {
	Name  string  `json:"book"`
	ISBN  string  `json:"isbn"`
	Price float64 `json:"price"`
}

func newResponseBook(book *entity.Book) *responseBook {
	return &responseBook{
		Name:  book.Name,
		ISBN:  book.ISBN,
		Price: book.Price,
	}
}

func GetBooksByID(app Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		repo, err := app.NewRepository(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				responseError{Message: err.Error()},
			)
			return
		}
		book := entity.Book{}
		if err := repo.GetBookByID(id, &book); err != nil {
			if errors.Is(err, NotFoundError) {
				ctx.AbortWithStatusJSON(
					http.StatusNotFound,
					responseError{Message: err.Error()},
				)
				return
			}
			if errors.Is(err, InvalidParameterError) {
				ctx.AbortWithStatusJSON(
					http.StatusBadRequest,
					responseError{Message: err.Error()},
				)
				return
			}
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				responseError{Message: err.Error()},
			)
			return
		}
		ctx.JSON(http.StatusOK, newResponseBook(&book))
	}
}

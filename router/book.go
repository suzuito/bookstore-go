package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/bookstore-go/entity"
)

type responseError struct {
	Message string `json:"message"`
}

type responseStatus struct {
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

type resopnseBooks struct {
	Books []*responseBook
}

func newResponseBooks(books *[]*entity.Book) *[]*responseBook {
	ret := []*responseBook{}
	for _, book := range *books {
		ret = append(ret, newResponseBook(book))
	}
	return &ret
}

func GetStatus(app Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responseStatus{Message: "ok"})
		return
	}
}

func GetBooks(app Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		repo, err := app.NewRepository(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				responseError{Message: err.Error()},
			)
			return
		}
		books := []*entity.Book{}
		if err := repo.GetBooks(&books); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				responseError{Message: err.Error()},
			)
			return
		}
		ctx.JSON(http.StatusOK, newResponseBooks(&books))
	}
}

func GetBooksByID(app Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responseError{
				Message: fmt.Sprintf("invalid id: %s", id),
			})
			return
		}
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

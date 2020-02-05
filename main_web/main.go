package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/bookstore-go/entity"
	"github.com/suzuito/bookstore-go/router"
	"github.com/suzuito/bookstore-go/router/repository"
)

type RepositoryInMemory struct{}

func NewRepositoryInMemory() *RepositoryInMemory {
	return &RepositoryInMemory{}
}

func (r *RepositoryInMemory) GetBooks(book *[]*entity.Book) error {
	return fmt.Errorf("Not impl")
}
func (r *RepositoryInMemory) GetBookByID(id string, book *entity.Book) error {
	return fmt.Errorf("Not impl")
}

type ApplicationImpl struct{}

func NewApplicationImpl() *ApplicationImpl {
	return &ApplicationImpl{}
}

func (a *ApplicationImpl) NewRepository(ctx context.Context) (repository.Repository, error) {
	return NewRepositoryInMemory(), nil
}

func main() {
	app := NewApplicationImpl()
	root := gin.Default()
	root.GET("/books", router.GetBooks(app))
	root.GET("/books/:id", router.GetBooksByID(app))
	if err := root.Run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

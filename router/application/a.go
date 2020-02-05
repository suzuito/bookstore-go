package application

import (
	"context"

	"github.com/suzuito/bookstore-go/router/repository"
)

type Application interface {
	NewRepository(ctx context.Context) (repository.Repository, error)
}

package router

import (
	"context"
)

type Application interface {
	NewRepository(ctx context.Context) (Repository, error)
}

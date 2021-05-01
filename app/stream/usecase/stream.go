package usecase

import "context"

type stream interface {
	createStream(ctx context.Context, ID, name, owner string) error
}

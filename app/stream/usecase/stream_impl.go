package usecase

import (
	"context"

	"github.com/leewei05/live-ddd/app/stream/adapter"
)

type streamUseCase struct {
	streamRepo adapter.Repository
}

func NewStreamUseCase(r adapter.Repository) *streamUseCase {
	return &streamUseCase{
		streamRepo: r,
	}
}

func (s *streamUseCase) createStream(ctx context.Context, ID, name, owner string) error {

	return nil
}

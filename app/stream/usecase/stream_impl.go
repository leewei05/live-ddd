package usecase

import (
	"context"

	adapter "github.com/leewei05/live-ddd/app/stream/adapter/repository"
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
	s.streamRepo.Create(context.Background(), ID, name, owner)

	return nil
}

package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/leewei05/live-ddd/app/stream/adapter/repository/mock_repo"
)

func TestCreateStream(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockRepository(ctrl)
	m.EXPECT().Create(
		context.Background(),
		"id_1",
		"Monday Stream!",
		"streamer_1",
	)

	m.Create(context.Background(), "id_1", "Monday Stream!", "streamer_1")
}

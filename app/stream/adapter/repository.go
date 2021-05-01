package adapter

import "context"

type Repository interface {
	// Create a new stream
	Create(ctx context.Context, ID, name, owner string)
	// Get all streams
	Get(ctx context.Context)
}

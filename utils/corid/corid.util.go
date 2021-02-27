package corid

import (
	"context"

	"github.com/google/uuid"
)

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key string

// k is the context key for the correlation id
const k = key("correlation-id")

// NewContext returns a new Context carrying correlation id
func NewContext(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, k, id)
}

// FromContext extracts the correlation id from ctx, if present.
func FromContext(ctx context.Context) (uuid.UUID, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the uuid.UUID type assertion returns ok=false for nil.
	corID, ok := ctx.Value(k).(uuid.UUID)
	return corID, ok
}

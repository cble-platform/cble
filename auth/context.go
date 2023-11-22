package auth

import (
	"context"
	"errors"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/internal/contexts"
)

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*ent.User, error) {
	raw, ok := ctx.Value(contexts.USER_CTX_KEY).(*ent.User)
	if ok {
		return raw, nil
	}
	return nil, errors.New("failed to get user from context")
}

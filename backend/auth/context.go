package auth

import (
	"context"
	"errors"

	"github.com/cble-platform/cble-backend/ent"
)

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*ent.User, error) {
	raw, ok := ctx.Value(USER_CTX_KEY).(*ent.User)
	if ok {
		return raw, nil
	}
	return nil, errors.New("failed to get user from context")
}

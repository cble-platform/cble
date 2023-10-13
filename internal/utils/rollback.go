package utils

import (
	"fmt"

	"github.com/cble-platform/backend/ent"
)

func RollbackWithErr(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

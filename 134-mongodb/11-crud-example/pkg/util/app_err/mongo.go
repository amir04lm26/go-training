package apperr

import (
	"errors"
	"fmt"
)

// Define custom error variables
var (
	ErrNoItemFoundToDelete = errors.New("no item found to delete")
	ErrNoItemFoundToUpdate = errors.New("no item found to update")
	ErrDeleteItemFailed    = fmt.Errorf("failed to delete the item: %w", ErrNoItemFoundToDelete)
	ErrUpdateItemFailed    = fmt.Errorf("failed to update the item: %w", ErrNoItemFoundToUpdate)
)

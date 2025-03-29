package ctxhelper

import (
	"context"
	"time"
)

var (
	CtxDefaultTimeout = 5 * time.Second
)

func CtxWithNormalTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), CtxDefaultTimeout)
}

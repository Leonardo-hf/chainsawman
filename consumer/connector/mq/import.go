package mq

import (
	msg "chainsawman/client/go"
	"context"
)

type ImportMq interface {
	ConsumeImportMsg(ctx context.Context, consumer string, handle func(ctx context.Context, m *msg.Msg) error) error
}

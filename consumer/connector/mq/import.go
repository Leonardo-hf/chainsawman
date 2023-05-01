package mq

import (
	"chainsawman/consumer/connector/msg"
	"context"
)

type ImportMq interface {
	ConsumeImportMsg(ctx context.Context, consumer string, handle func(ctx context.Context, m *msg.Msg) error) error
}

package mq

import (
	"chainsawman/consumer/connector/model"
	"context"
)

type ImportMq interface {
	ConsumeImportMsg(ctx context.Context, consumer string, handle func(ctx context.Context, m *model.Msg) error) error
}

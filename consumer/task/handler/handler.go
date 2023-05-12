package handler

import "chainsawman/consumer/task/model"

type Handler interface {
	Handle(task *model.KVTask) (string, error)
}

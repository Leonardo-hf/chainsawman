package handler

type Handler interface {
	Handle(params string, taskID int64) (string, error)
}

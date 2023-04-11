package handler

type Handler interface {
	Handle(params string) (string, error)
}

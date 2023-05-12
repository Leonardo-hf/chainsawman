package handler

import msg "chainsawman/client/go"

type Handler interface {
	Handle(m *msg.Msg) error
}

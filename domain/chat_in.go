package domain

type ChatIn interface {
	UserChat() (*UserChat, error)
}

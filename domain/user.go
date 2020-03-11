package domain

type User struct {
	ID            int
	Name          string
	SocialService []SocialService
}

type UserRepository interface {
	Save()
	ReadAll()
	Raed()
	Delete()
}

type UserContent struct {
	ID       int
	UserID   int
	SocialID int
	Content  interface{}
}

type UserContentRepository interface {
	Save()
	Read()
	ReadAll()
	Delete()
}

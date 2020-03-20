package domain

type User struct {
	ID            int
	Name          string
	SocialService []SocialService
}

type UserRepository interface {
	Save(User) error
	Read() (User, error)
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

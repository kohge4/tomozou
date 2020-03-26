package domain

type UserRepository interface {
	Save(User) (int, error)
	ReadByID(id int) (User, error)
	ReadBySocialID(socialID string) ([]User, error)
	ReadAll() []User
}

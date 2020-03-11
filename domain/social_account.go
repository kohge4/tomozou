package domain

type SocialService struct {
	SocialID    string
	ServiceName string
}

type SocialAccount interface {
	ToUser() (User, error)
	Content() (SocialAccountContent, error)

	SignIn()
	SignUp()
	SignOut(User) error
	DeleteUser(User) error
}

type SocialAccountContent interface {
	Save()
	//ReadAll()
	Read()
	//Delete()
	ToUserContent() (UserContent, error)
}

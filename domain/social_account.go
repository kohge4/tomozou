package domain

type SocialService struct {
	SocialID    string
	ServiceName string
}

type SocialAccount interface {
	User() (User, error)
	Content() (SocialAccountContent, error)

	SignIn() error
	SignUp() error
	SignOut(User) error
	//DeleteUser(User) error
}

type SocialAccountContent interface {
	Save()
	//ReadAll()
	Read()
	//Delete()
	ToUserContent() (UserContent, error)
}

type SocialAccountRepository interface{}

package port

import "tomozou/domain"

type WebServiceAccount interface {
	User() (domain.User, error)
	Content() (ContentRepository, error)

	SignIn() error
	SignUp() error
	SignOut(domain.User) error
	//DeleteUser(User) error
}

type ContentRepository interface {
	Save(userID int) error
	//ReadAll()
	//Read()
	//Delete()
	//UserContent() ([]UserContent, error)
}

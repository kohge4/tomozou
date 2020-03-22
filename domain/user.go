package domain

type User struct {
	ID       int    `gorm:"not null;AUTO_INCREMENT" json:"id"`
	SocialID string `gorm:"not null" json:"social_id"`
	Name     string `gorm:"not null" json:"name"`
	Auth     string `gorm:"not null" json:"auth"`
	Image    string `gorm:"column:image" json:"image"`
}

func NewUser(socialID string, name string, auth string, image string) User {
	return User{
		SocialID: socialID,
		Name:     name,
		Auth:     auth,
		Image:    image,
	}
}

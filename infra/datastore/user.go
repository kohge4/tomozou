package datastore

import "github.com/jinzhu/gorm"

type UserDBRepository struct {
	DB *gorm.DB
}

func NewUserDBRepository() UserDBRepository {
	db, _ := GormConn()
	db.AutoMigrate(&User{})
	uRepo := UserDBRepository{
		DB: db,
	}
	return uRepo
}

type User struct {
	ID       int    `gorm:"not null;AUTO_INCREMENT" json:"id"`
	SocialID string `gorm:"not null;AUTO_INCREMENT" json:"social_id"`
	Name     string `gorm:"not null" json:"name"`
	Auth     string `gorm:"not null" json:"auth"`
	Image    string `gorm:"column:image" json:"image"`

	/*
		ScreenName  string     `gorm:"not null" json:"screen_name"`
		Name        string     `gorm:"not null" json:"name"`
		URL         string     `gorm:"not null" json:"url"`
		Description string     `gorm:"null" json:"description"`
		IsSignedIn  bool       `gorm:"not null" json:"is_signed_in"`
		CreatedAt   time.Time  `gorm:"null" json:"create_at"`
		UpdatedAt   time.Time  `gorm:"null" json:"update_at"`
		DeletedAt   *time.Time `gorm:"null" json:"-"`
	*/
}

func (repo *UserDBRepository) Save(user User) error {
	// 本来は domain.User => infra 用の User に変換する interface を用いたい
	repo.DB.Create(&user)
	return nil
}

func (repo *UserDBRepository) ReadAll() []User {
	users := []User{}
	repo.DB.Find(&users)
	return users
}

//func (repo *UserDBRepository) Read

package datastore

import (
	"time"

	"github.com/jinzhu/gorm"
)

// spotify に 関する user の 情報を 表示

type UserSpotifyDBRepository struct {
	DB *gorm.DB
}

func NewUserSpotifyDBRepository() UserSpotifyDBRepository {
	db, _ := GormConn()
	db.AutoMigrate(&UserSpotify{})
	uRepo := UserSpotifyDBRepository{
		DB: db,
	}
	return uRepo
}

// Login 時 に accesstoken とかを保存する
type UserSpotify struct {
	USID         int    `gorm:"column:usID;not null;AUTO_INCREMENT" json:"us_id"`
	UserID       string `gorm:"column:userID;not null" json:"user_id"`
	UserName     string `gorm:"column:userName;not null" json:"user_name"`
	SocialID     string `gorm:"column:socialID;not null" json:"social_id"`
	AccessToken  string `gorm:"column:accessToken;not null" json:"accesstoken"`
	RefreshToken string `gorm:"column:refreshToken;not null" json:"accesstoken"`
	Playlist     string `gorm:"column:playlist" json:"playlist"`
	Favorite     string `gorm:"column:favorite" json:"favorite"`
}

func (repo *UserSpotifyDBRepository) Save(user UserSpotify) error {
	// 本来は domain.User => infra 用の User に変換する interface を用いたい
	repo.DB.Create(&user)
	return nil
}

func (repo *UserSpotifyDBRepository) ReadAll() []UserSpotify {
	users := []UserSpotify{}
	repo.DB.Find(&users)
	return users
}

type UserSpotifyPlaylist struct {
	ID          int
	UserID      string
	PlaylistURL string
	Comment     string
}

type UserSpotifyArtist struct {
	UserID    string
	CreatedAt time.Time
	//Timerange string
	// short か long か
	ArtistType string

	SpArtistID    string
	SpMusicURL    string
	SpArtistImage string
}

type UserSpotifyTrack struct {
	UserID    string
	CreatedAt time.Time
	TrackType string
	// Nowplaying か何か
	SpTrackID    string
	SpTrackURL   string
	SpTrackImage string
}

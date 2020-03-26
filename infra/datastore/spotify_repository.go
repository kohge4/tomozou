package datastore

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

// SpotifyHanlder が 構造体 に もつ リポジトリ
type SpotifyItemDBRepository struct {
	DB *gorm.DB
}

func NewSpotifyItemDBRepository(db *gorm.DB) domain.ItemRepository {
	return &SpotifyItemDBRepository{
		DB: db,
	}
}

// UserApplication で 外から使用する ==> 大元の リポジトリを 外から使用する方針の方が綺麗
//　いらない説
func (repo *SpotifyItemDBRepository) ReadItemByUser(userID int) (interface{}, error) {
	var artists []domain.Artist
	repo.DB.Find(&artists)
	return artists, nil
}

func (repo *SpotifyItemDBRepository) ReadTagByUser(userID int) (interface{}, error) {
	var tag []domain.UserArtistTag
	repo.DB.Find(&tag)
	return tag, nil
}

// 以下は SpotifyHandler から 保存するときガンガン使用する
func (repo *SpotifyItemDBRepository) SaveArtist(artist domain.Artist) (int, error) {
	repo.DB.Create(&artist)
	return artist.ID, nil
}

func (repo *SpotifyItemDBRepository) SaveTrack(domain.Track) (int, error) {
	return 0, nil
}

func (repo *SpotifyItemDBRepository) SaveUserArtistTag(tag domain.UserArtistTag) error {
	repo.DB.Create(&tag)
	return nil
}

func (repo *SpotifyItemDBRepository) ReadArtistBySocialID(socialID string) (*domain.Artist, error) {
	//repo.DB.Find()
	return nil, nil
}

func (repo *SpotifyItemDBRepository) ReadUserArtistTagByUserID(userID int) (interface{}, error) {
	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("user_id = ?", userID).Find(&userArtistTags)
	return userArtistTags, nil
}

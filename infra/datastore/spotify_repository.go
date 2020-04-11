package datastore

import (
	"fmt"
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

func (repo *SpotifyItemDBRepository) SaveUserArtistTag(tag domain.UserArtistTag) error {
	repo.DB.Create(&tag)
	return nil
}

func (repo *SpotifyItemDBRepository) SaveTrack(track domain.Track) (int, error) {
	repo.DB.Create(&track)
	return track.ID, nil
}

func (repo *SpotifyItemDBRepository) SaveUserTrackTag(tag domain.UserTrackTag) error {
	repo.DB.Create(&tag)
	return nil
}

func (repo *SpotifyItemDBRepository) ReadArtistBySocialID(socialID string) (*domain.Artist, error) {
	artists := []*domain.Artist{}
	repo.DB.Where("social_id = ?", socialID).Find(&artists)
	if len(artists) == 0 {
		return nil, nil
	}
	if len(artists) == 1 {
		return artists[0], nil
	}
	fmt.Println("DUPLICATED ARTIST")
	return artists[0], nil
}

func (repo *SpotifyItemDBRepository) ReadUserArtistTagByUserID(userID int) (interface{}, error) {
	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("user_id = ?", userID).Find(&userArtistTags)
	return userArtistTags, nil
}

func (repo *SpotifyItemDBRepository) ReadUserIDByArtistID(artistID int) ([]int, error) {
	var users []int
	flag := 0

	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("artist_id = ?", artistID).Find(&userArtistTags)

	// 並び替えた後
	for _, tag := range userArtistTags {
		if flag != tag.UserID {
			users = append(users, tag.UserID)
			flag = tag.UserID
		}
	}
	return users, nil
}

func (repo *SpotifyItemDBRepository) ReadUserIDByArtistName(artistName string) ([]int, error) {
	var users []int
	flag := 0

	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("artist_name = ?", artistName).Find(&userArtistTags)

	// 並び替えた後
	for _, tag := range userArtistTags {
		if flag != tag.UserID {
			users = append(users, tag.UserID)
			flag = tag.UserID
		}
	}
	return users, nil
}

func (repo *SpotifyItemDBRepository) ReadUserArtistTagByTagID(tagID int) (interface{}, error) {
	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("tag_id = ?", tagID).Find(&userArtistTags)
	return userArtistTags, nil
}

func (repo *SpotifyItemDBRepository) ReadUserTrackTagByUserID(userID int) ([]domain.UserTrackTag, error) {
	// nowplaying の 表示用
	userTrackTags := []domain.UserTrackTag{}
	repo.DB.Where("user_id = ?", userID).Find(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}

func (repo *SpotifyItemDBRepository) DeleteAllUserArtistTagsByUserID(userID int) error {
	tag := domain.UserArtistTag{}
	repo.DB.Where("user_id LIKE ?", userID).Delete(&tag)
	return nil
}

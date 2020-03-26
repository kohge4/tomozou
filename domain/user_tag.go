package domain

type UserArtistTag struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID   int    `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID int    `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName  string `gorm:"column:tag_name;not null" json:"tag_name"`

	ArtistName string `gorm:"column:artist_name" json:"artist_name"`
	URL        string `gorm:"column:url" json:"url"`
	Image      string `gorm:"column:image" json:"image"`
}

func NewUserArtistTag(userID int, artistID int, tagName string) UserArtistTag {
	return UserArtistTag{
		UserID:   userID,
		ArtistID: artistID,
		TagName:  tagName,
	}
}

type UserTrackTag struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID  int    `gorm:"column:user_id;not null" json:"user_id"`
	TrackID int    `gorm:"column:user_id;not null" json:"track_id"`
	TagName string `gorm:"column:tag_name;not null" json:"tag_name"`

	ArtistName string `gorm:"column:artist_name" json:"artist_name"`
	TrackName  string `gorm:"column:track_name" json:"track_name"`
	URL        string `gorm:"column:url" json:"url"`
}

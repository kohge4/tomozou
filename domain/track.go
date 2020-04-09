package domain

type Track struct {
	ID       int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	SocialID string `gorm:"column:name;not null" json:"social_id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	TrackURL string `gorm:"column:track_url;not null" json:"track_url"`

	ArtistName string `gorm:"column:arttist_name;not null" json:"artist_name"`
	ArtistID   int    `gorm:"column:arttist_id;not null" json:"artist_id"`
}

func (t *Track) UserTrackTag(userID int) *UserTrackTag {
	return NewUserTrackTag(t, userID)
}

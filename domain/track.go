package domain

type Track struct {
	ID       int
	SocialID string
	Name     string
	TrackURL string

	ArtistName string
	ArtistID   int
}

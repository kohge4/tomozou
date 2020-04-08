package webservice

import (
	"tomozou/domain"

	"github.com/kohge4/spotify"
)

func SimpleTrackToTrack(simpleTrack *spotify.SimpleTrack) *domain.Track {
	return &domain.Track{
		SocialID:   simpleTrack.ID.String(),
		Name:       simpleTrack.Name,
		TrackURL:   simpleTrack.ExternalURLs["spotify"],
		ArtistName: simpleTrack.Artists[0].Name,
		ArtistID:   simpleTrack.Artists[0].ID.String(),
	}
}

func ArtistsFromTrack(simpleTrack *spotify.SimpleTrack) []*domain.Artist {
	var artistList []*domain.Artist

	for _, artist := range simpleTrack.Artists {
		artistIn := &domain.Artist{
			Name: artist.Name,
		}
		artistList = append(artistList, artistIn)
	}
	return artistList
}

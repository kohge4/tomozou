package webservice

import (
	"tomozou/domain"

	"github.com/kohge4/spotify"
)

func SimpleTrackToTrack(h *SpotifyHandler, simpleTrack *spotify.SimpleTrack) *domain.Track {
	return &domain.Track{
		SocialID:   simpleTrack.ID.String(),
		Name:       simpleTrack.Name,
		ArtistName: simpleTrack.Artists[0].Name,
		ArtistID:   ArtistIDFromSpotifyID(h, simpleTrack.Artists[0].ID.String()),
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

func ArtistIDFromSpotifyID(h *SpotifyHandler, spotifyID string) int {
	// 最初に そのアーティストが存在していなかったら 保存の処理をする
	artist, _ := h.SpotifyRepository.ReadArtistBySocialID(spotifyID)
	return artist.ID
}

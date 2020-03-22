package webservice

import (
	"fmt"
	"tomozou/domain"

	"github.com/kohge4/spotify"
)

func (h *SpotifyHandler) saveTopArtists(userID int) error {
	timerange := "long"
	opt := &spotify.Options{
		Timerange: &timerange,
	}

	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	fmt.Println(results)
	fmt.Printf("%T", results)
	if err != nil {
		return err
	}
	for _, result := range results.Items {
		var artist *domain.Artist
		artist, _ = h.SpotifyRepository.ReadArtistBySocialID(result.ID)
		if artist == nil {
			artist = &domain.Artist{
				Name:     result.Name,
				SocialID: result.ID,
				Image:    result.Images[0].URL,
			}
			artist.ID, err = h.SpotifyRepository.SaveArtist(*artist)
			if err != nil {
				return err
			}
		}
		tag := domain.UserArtistTag{
			UserID:   userID,
			ArtistID: artist.ID,
			TagName:  "top_artist",
		}
		h.SpotifyRepository.SaveUserArtistTag(tag)
	}
	return nil
}

func (h *SpotifyHandler) saveRecentlyFavoriteArtists(userID int) error {
	timerange := "short"
	opt := &spotify.Options{
		Timerange: &timerange,
	}

	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	if err != nil {
		return err
	}
	for _, result := range results.Items {
		var artist *domain.Artist
		artist, _ = h.SpotifyRepository.ReadArtistBySocialID(result.ID)
		if artist == nil {
			artist = &domain.Artist{
				Name:     result.Name,
				SocialID: result.ID,
				Image:    result.Images[0].URL,
			}
			artist.ID, err = h.SpotifyRepository.SaveArtist(*artist)
			if err != nil {
				return err
			}
		}
		tag := domain.UserArtistTag{
			UserID:   userID,
			ArtistID: artist.ID,
			TagName:  "recently_favorite_artist",
		}
		h.SpotifyRepository.SaveUserArtistTag(tag)
	}
	return nil
}

func (h *SpotifyHandler) saveRecentlyPlayedTracks(useraID int) error {
	return nil
}

func (h *SpotifyHandler) saveTopTracks() error {
	return nil
}

func (h *SpotifyHandler) saveNowPlayingTrack() error {
	return nil
}

func (h *SpotifyHandler) checkDupulicateArtist(socialID string) bool {
	artist, _ := h.SpotifyRepository.ReadArtistBySocialID(socialID)
	if artist == nil {
		return false
	}
	return true
}

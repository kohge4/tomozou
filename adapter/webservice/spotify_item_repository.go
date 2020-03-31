package webservice

import (
	"tomozou/domain"

	"github.com/kohge4/spotify"
)

func (h *SpotifyHandler) saveTopArtists(userID int) error {
	timerange := "long"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	if err != nil {
		return err
	}
	for _, result := range results.Items {
		var artist *domain.Artist

		artist, _ = h.SpotifyRepository.ReadArtistBySocialID(result.ID)
		println("CCCHHEEC")
		println(artist)
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
			UserID:     userID,
			ArtistID:   artist.ID,
			TagName:    "top_artist",
			ArtistName: result.Name,
			URL:        result.ExternalUrls.Spotify,
			Image:      result.Images[0].URL,
		}
		h.SpotifyRepository.SaveUserArtistTag(tag)
	}
	return nil
}

func (h *SpotifyHandler) saveRecentlyFavoriteArtists(userID int) error {
	timerange := "short"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
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
			UserID:     userID,
			ArtistID:   artist.ID,
			TagName:    "recently_favorite_artist",
			ArtistName: result.Name,
			URL:        result.ExternalUrls.Spotify,
			Image:      result.Images[0].URL,
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

package webservice

import (
	"fmt"
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

func (h *SpotifyHandler) deleteUserArtistInfo(userID int) error {
	//h.SpotifyRepository.DeleteAllArtitByUserID(userID)
	err := h.SpotifyRepository.DeleteAllUserArtistTagsByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}

/*
func (h *SpotifyHandler) saveRecentlyPlayedTracks(useraID int) error {
	// あとで
	//timerange := "short"
	limit := 5
	opt := &spotify.RecentlyPlayedOptions{
		//Timerange: &timerange,
		Limit: limit,
	}

	results, err := h.Client.PlayerRecentlyPlayedOpt(opt)
	if err != nil {
		return err
	}
	/*
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
*/

func (h *SpotifyHandler) saveTopTracks(userID int) error {
	/*
		trackのデータを取得
		artist データがなければ追加 => artist保存
		trackの保存データに変換
	*/
	timerange := "short"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	results, err := h.Client.GetUserTopTracks2Opt(opt)
	if err != nil {
		return err
	}
	println("TTTTTrack")
	fmt.Println(results.Items)
	/*
		for _, result := range results.Items {
			var artist *domain.Artist

			artist, _ = h.SpotifyRepository.ReadArtistBySocialID(result.ID)
			if artist == nil {
				artist = &domain.Artist{
					Name:     result.Name,
					SocialID: result.ID,
					Image:    result.Album.Images[0].URL,
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
				Image:      result.Album.Images[0].URL,
			}
			h.SpotifyRepository.SaveUserArtistTag(tag)
		}
	*/
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

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
	for _, result := range results.Items {
		var artistIn *domain.Artist
		var trackIn *domain.Track

		artists := result.Album.Artists

		artistIn, _ = h.SpotifyRepository.ReadArtistBySocialID(artists[0].ID)
		if artistIn == nil {
			artistIn = &domain.Artist{
				Name:     result.Artists[0].Name,
				SocialID: result.Artists[0].ID,
				Image:    "",
			}
			artistIn.ID, err = h.SpotifyRepository.SaveArtist(*artistIn)
			if err != nil {
				return err
			}
		}

		trackIn = &domain.Track{
			Name: result.Name,
			// TrackURL ではなくsocialID で url作る方針
			SocialID:   result.ID,
			ArtistName: artistIn.Name,
			ArtistID:   artistIn.ID,
		}
		_, err := h.SpotifyRepository.SaveTrack(*trackIn)
		if err != nil {
			return err
		}
		// 複数の arthist が 携わるトラックの場合の処理
		// corrywong の cosmic sans を 聞いて nowplaying の処理とともに 確認
	}
	// UserTrack Tag の 保存も必要 そのための userID
	return nil
}

func (h *SpotifyHandler) saveNowPlayingTrack() error {
	var track *spotify.SimpleTrack

	var artistIn *domain.Artist
	var trackIn *domain.Track

	nowPlaying, err := h.Client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}

	if nowPlaying.Item != nil {
		track = nowPlaying.Item.ToSimpleTrack()
	} else {
		recentlyPlaying, err := h.Client.PlayerRecentlyPlayed()
		if err != nil {
			return err
		}
		track = &recentlyPlaying[0].Track
	}
	artists := track.Artists

	artistIn, _ = h.SpotifyRepository.ReadArtistBySocialID(artists[0].ID.String())
	if artistIn == nil {
		artistIn = &domain.Artist{
			Name:     track.Artists[0].Name,
			SocialID: track.Artists[0].ID.String(),
			Image:    "",
		}
		artistIn.ID, err = h.SpotifyRepository.SaveArtist(*artistIn)
		if err != nil {
			return err
		}
	}
	trackIn = &domain.Track{
		Name: track.Name,
		// TrackURL ではなくsocialID で url作る方針
		SocialID:   track.ID.String(),
		ArtistName: artistIn.Name,
		ArtistID:   artistIn.ID,
	}
	// Track 保存に関する処理
	// SimpleTrack を変換 => artist を保存 => tagとして track に持たせる
	h.SpotifyRepository.SaveTrack(*trackIn)
	// この後に Tag の 保存も必要  そのための userID
	return nil
}

func (h *SpotifyHandler) checkDupulicateArtist(socialID string) bool {
	artist, _ := h.SpotifyRepository.ReadArtistBySocialID(socialID)
	if artist == nil {
		return false
	}
	return true
}

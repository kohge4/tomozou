package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *SpotifyHandler) SaveUserSpotifyDataLogin(c *gin.Context) {
	// 他で やったからいらないね
}

func (h *SpotifyHandler) SaveUserCurrentlyPlaylist(c *gin.Context) {
	// 正規化 するから テーブルは複数かもね
}

func (h *SpotifyHandler) SaveUserRecentlyPlayedTracks(c *gin.Context) {

}

func (h *SpotifyHandler) SaveUserFavoriteArtists(c *gin.Context) {
	// 時間ごとに 表示できるにしたい

}

func (h *SpotifyHandler) SaveUserNowPlaying(c *gin.Context) {

}

func (h *SpotifyHandler) DisplayFavoriteArtist(c *gin.Context) {
	// ある程度 表示して 保存する
}

func (h *SpotifyHandler) DisplayRecentlyPlayingAritst(c *gin.Context) {}

// timerange short_term
func (h *SpotifyHandler) MeTopArtistShortTerm(c *gin.Context) {
	result, err := h.Client.CurrentUsersAlbumsOpt(nil)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

// timerange long_term
func (h *SpotifyHandler) MeTopArtistLongTerm(c *gin.Context) {
	result, err := h.Client.CurrentUsersAlbumsOpt(nil)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

// 最近再生した 曲を 5曲ほど保存
func (h *SpotifyHandler) MeRecentlyPlaying(c *gin.Context) {
	result, err := h.Client.PlayerRecentlyPlayed()
	// RecentlyPlayedOption で 表示数を制限できる
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

// 今再生している曲を表示
func (h *SpotifyHandler) MeNowPlaying(c *gin.Context) {
	result, err := h.Client.PlayerCurrentlyPlaying()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

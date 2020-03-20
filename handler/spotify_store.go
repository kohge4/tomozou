package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kohge4/spotify"
)

func (h *SpotifyHandler) SaveUserSpotifyDataLogin(c *gin.Context) {
	// 他で やったからいらないね
}

func (h *SpotifyHandler) SaveUserCurrentlyPlaylist(c *gin.Context) {
	// 正規化 するから テーブルは複数かもね
}

func (h *SpotifyHandler) SaveUserRecentlyPlayedTracks(c *gin.Context) {
	//recentlyPlayedTracks := h.Client.
}

/*
func (h *SpotifyHandler) saveArtists(c *gin.Context, data interface{}) {
	resp, ok := data.(*spotify.CurrentUserArtistResponse)
	if ok == false {
		fmt.Println(resp)
		return
	}

	//var artist datastore.Artist
	//var userArtistTag datastore.UserArtistTag
	tomozouID, _ := c.Get("tomozou-id")
	userID, _ := tomozouID.(int)

	for _, item := range resp.Items {
		// item.Images[0] が ないときの 条件分岐欲しい

		// socialID で Artist があるか 探す, なければ追加
		artist := datastore.NewArtist(item.Name, item.ID, item.Images[0].URL)

		userArtistTag := datastore.NewUserArtistTag(userID, 1, "short_term")

		// spotifyHandler に SpotifyDBRepository を　作成
		// => そこから メソッドを切り分けていくプラン
	}

}
*/

func (h *SpotifyHandler) SaveUserFavoriteArtists(c *gin.Context) {
	var timerange string
	var opt spotify.Options

	timerange = "short_term"
	opt.Timerange = &timerange
	favorite, err := h.Client.CurrentUsersTopArtistsOpt(&opt)
	if err != nil {
		fmt.Println(err)
	}

	timerange = "long_term"
	opt.Timerange = &timerange
	recentlyFavorite, err := h.Client.CurrentUsersTopArtistsOpt(&opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(recentlyFavorite)

	c.JSON(200, favorite)
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

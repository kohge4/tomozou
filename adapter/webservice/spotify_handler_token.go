package webservice

import (
	"fmt"
	"tomozou/domain"

	"golang.org/x/oauth2"
)

func (h *SpotifyHandler) saveUserToken(userID int) error {
	spToken, err := h.Client.Token()
	if err != nil {
		return err
	}
	token := newUserToken(userID, "spotify", spToken)
	h.DB.Create(token)
	return nil
}

func (h *SpotifyHandler) updateUserToken(userID int) error {
	// 怪しい. あってるか確認してない
	spToken, err := h.Client.Token()
	if err != nil {
		return err
	}
	userToken := &domain.UserToken{}
	h.DB.Where("user_id = ?", userID).First(userToken)
	userToken = updateUserToken(userToken, spToken)
	h.DB.Save(userToken)
	return nil
}

func (h *SpotifyHandler) readUserToken(token *domain.UserToken) {
	h.DB.First(token)
	fmt.Println(token)
}

func newUserToken(userID int, authType string, token *oauth2.Token) *domain.UserToken {
	return &domain.UserToken{
		UserID:       userID,
		AuthType:     authType,
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
}

func updateUserToken(userToken *domain.UserToken, token *oauth2.Token) *domain.UserToken {
	userToken.AccessToken = token.AccessToken
	userToken.RefreshToken = token.RefreshToken
	userToken.Expiry = token.Expiry
	return userToken
}

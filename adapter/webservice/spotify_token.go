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
	if !h.DB.HasTable(&domain.UserToken{}) {
		h.DB.CreateTable(&domain.UserToken{})
	}
	h.DB.Create(token)
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

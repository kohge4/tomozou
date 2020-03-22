package domain

import "time"

type UserToken struct {
	// 基本的に ID と UsrID は 同じ値である必要がある
	ID       int    `gorm:"not null" json:"id"`
	UserID   int    `gorm:"not null" json:"user_id"`
	AuthType string `gorm:"not null" json:"auth_type"`

	AccessToken  string    `gorm:"not null" json:"access_token"`
	TokenType    string    `gorm:"not null" json:"token_type,omitempty"`
	RefreshToken string    `gorm:"not null" json:"refresh_token,omitempty"`
	Expiry       time.Time `gorm:"not null" json:"expiry,omitempty"`
}

type UserSubToken struct {
	// 二種類目以降の連携に用いるはず
	ID       int    `json:"id"`
	TokenID  int    `json:"token_id"`
	UserID   int    `json:"user_id"`
	AuthType string `json:"auth_type"`

	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

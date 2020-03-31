package usecase

import (
	"tomozou/domain"
)

//
/*
必要な機能を書き出す

LoginしたUser
	UserProfile の編集
	ログイン状態 の変更
	再連携social service の 情報 表示

User 全体

application として必要な機能
*/

/*
認証の 抽象化 難しい
*/

/*
認証以外の Usecase
User の 情報を 保存(Login 情報に基づき)
User の 情報を表示(DB から)

*/

type UserProfileApplication struct {
	UserRepository domain.UserRepository
	ItemRepository domain.ItemRepository

	// spotify 関連の情報は まとめて 保存する (やりとりはない)
	WebServiceAccount domain.WebServiceAccount
}

func NewUserProfileApplication(uR domain.UserRepository, iR domain.ItemRepository) *UserProfileApplication {
	return &UserProfileApplication{
		UserRepository: uR,
		ItemRepository: iR,
	}
}

func (u *UserProfileApplication) RegistryUser() error {
	// アカウントを登録して User 情報を保存する
	user, err := u.WebServiceAccount.User()
	id, err := u.UserRepository.Save(*user)
	if err != nil {
		return err
	}
	err = u.WebServiceAccount.SaveUserItem(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserProfileApplication) ReRegistiryUser(id int) error {
	// 任意の User の アカウントを 再連携して, 情報を更新する
	// Token を 外からひっぱてくる処理をかく
	user, err := u.UserRepository.ReadByID(id)
	if err != nil {
		return err
	}
	err = u.WebServiceAccount.Link(user)
	if err != nil {
		return err
	}
	err = u.WebServiceAccount.SaveUserItem(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserProfileApplication) Me(id int) (interface{}, error) {
	me, err := u.UserRepository.ReadByID(id)
	if err != nil {
		return nil, err
	}
	return me, nil
}

func (u *UserProfileApplication) DisplayMe(id int) (interface{}, error) {
	artists, err := u.ItemRepository.ReadItemByUser(id)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (u *UserProfileApplication) MyArtistTag(id int) (interface{}, error) {
	artistTags, err := u.ItemRepository.ReadUserArtistTagByUserID(id)
	if err != nil {
		return nil, err
	}
	return artistTags, nil
}

func (u *UserProfileApplication) DisplayContent(id int) (interface{}, error) {
	// Controller で id を token から 持ってくる
	item, err := u.ItemRepository.ReadItemByUser(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (u UserProfileApplication) MyUserArtistTag(id int) (interface{}, error) {
	tags, err := u.ItemRepository.ReadUserArtistTagByUserID(id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (u UserProfileApplication) DisplayUsersByArtist(artistID int) (interface{}, error) {
	var users []domain.User
	var user domain.User

	// Userが tag が新しい順にソートされてる
	userIDs, err := u.ItemRepository.ReadUserIDByArtistID(artistID)
	for _, userID := range userIDs {
		user, err = u.UserRepository.ReadByID(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

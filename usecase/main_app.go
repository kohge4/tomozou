package usecase

import (
	"fmt"
	"tomozou/domain"
)

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

func (u *UserProfileApplication) RegistryUser() (*domain.User, error) {
	// アカウントを登録して User 情報を保存する
	user, err := u.WebServiceAccount.User()
	id, err := u.UserRepository.Save(*user)
	user.ID = id
	if err != nil {
		return nil, err
	}
	err = u.WebServiceAccount.SaveUserItem(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserProfileApplication) CheckExistingUser() (*domain.User, error) {
	// すでに 連携したことのあるユーザーの場合 他の処理にする
	user, err := u.WebServiceAccount.User()
	if err != nil {
		return nil, err
	}
	socialUsers, err := u.UserRepository.ReadBySocialID(user.SocialID)
	if err != nil {
		return nil, err
	}
	if len(socialUsers) == 0 {
		return nil, nil
	}
	return &socialUsers[0], nil
}

func (u *UserProfileApplication) UpdateUser(id int) error {
	// 任意の User の アカウントを 再連携して, 情報を更新する
	_, err := u.UserRepository.Update(id)
	if err != nil {
		// 最終更新日みたいなのを登録できるようにしたい
		return err
	}
	/*
		UpdateUserItem は 過去の情報を多少保持しておいたほうがいいかも
		回数とか記録しときたい気もする
	*/
	err = u.WebServiceAccount.UpdateUserItem(id)
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

func (u UserProfileApplication) MyUserArtistTag(id int) (interface{}, error) {
	tags, err := u.ItemRepository.ReadUserArtistTagByUserID(id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (u UserProfileApplication) MyUserTrackTag(id int) (interface{}, error) {
	// nowplaying を表示する用
	trackTag, err := u.ItemRepository.ReadUserTrackTagByUserID(id)
	if err != nil {
		return nil, err
	}
	return trackTag, nil
}

func (u UserProfileApplication) MyNowPlayingUserTrackTag(id int) (*domain.UserTrackTag, error) {
	// nowplaying を表示する用
	trackTags, err := u.ItemRepository.ReadUserTrackTagByUserID(id)
	if err != nil {
		return nil, err
	}
	if len(trackTags) == 0 {
		return nil, fmt.Errorf("nil error")
	}
	return &trackTags[len(trackTags)-1], nil
}

func (u *UserProfileApplication) FetchNowPlayng(id int) (interface{}, error) {
	// nowplayng を　外部から読み取った上で表示
	_, err := u.UserRepository.Update(id)
	if err != nil {
		// 最終更新日みたいなのを登録できるようにしたい
		return nil, err
	}
	// この options は domain に 厳密に型を用意してやった方がいいかも
	err = u.WebServiceAccount.UpdateUserItemOpt(id, "nowplaying")
	if err != nil {
		return nil, err
	}
	trackTag, err := u.ItemRepository.ReadUserTrackTagByUserID(id)
	if err != nil {
		return nil, err
	}
	return trackTag, nil
}

// UserID から その UserIDの もつ artistID を　全部 検索する
func (u UserProfileApplication) DisplayUsersByArtistID(artistID int) (interface{}, error) {
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

func (u UserProfileApplication) DisplayUsersByArtistName(artistName string) (interface{}, error) {
	var users []domain.User
	var user domain.User

	// Userが tag が新しい順にソートされてる
	userIDs, err := u.ItemRepository.ReadUserIDByArtistName(artistName)
	for _, userID := range userIDs {
		user, err = u.UserRepository.ReadByID(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

/*
// 実装開始　追加機能
func (u UserProfileApplication) AddUserArtistTagComment(tagID int, comment string) (interface{}, error) {
	var users []domain.User
	var user domain.User

	// Userが tag が新しい順にソートされてる
	userIDs, err := u.ItemRepository.ReadUserIDByArtistName(tagID)
	for _, userID := range userIDs {}
		user, err = u.UserRepository.ReadByID(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
*/

func (u UserProfileApplication) AddUserArtistTagComment(tagID int, comment string) (interface{}, error) {
	var users []domain.User

	return users, nil
}

func (u UserProfileApplication) DisplayNumberOfUserArtistTag() {}

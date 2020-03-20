package usecase

import "tomozou/domain"

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

type UserApplication struct {
	UserRepository domain.UserRepository
	SocialAccount  domain.SocialAccount
}

func (u UserApplication) SignUp() error {
	err := u.SocialAccount.SignIn()
	if err != nil {
		return err
	}

	user, err := u.SocialAccount.User()
	if err != nil {
		return err
	}

	err = u.UserRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserApplication) SignIn() error {
	_, err := u.UserRepository.Read()
	if err != nil {
		return nil
	}
	return nil
}

func (u UserApplication) SignOut() {

}

func (u UserApplication) DisplayContent() {

}

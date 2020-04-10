package domain

type WebServiceAccount interface {
	// ログイン内容に基づいて Social アカウントの情報を domain.User 形式に 出力
	User() (*User, error)
	//Link(User) error

	// SpotifyHandler.SaveUserItem で 必要な情報を 全部保存
	SaveUserItem(userID int) error

	// 再連携時の Item 周りの処理
	UpdateUserItem(userID int) error
	UpdateUserItemOpt(userID int, opt string) error
}

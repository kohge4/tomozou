package usecase 

/*

ログイン 構成
1 未 twitter/spotify で 新規登録する
2 未 twitter/spotify で ログインする
3 済 twitter/spotify を 連携する
4 済 twitter/spotify を　再連携する

1 JWTの accesstoken も refreshtoken も　ない状態 ==> cookie にない
2 JWT を cookie から　削除している or 期限切れ であるが 過去にアクセス(登録)したことある状態
      == twitter/spotify の ID で検索=> JWT 取得 (このアカウントで登録したことはないという画面を作る)
3 JWT に scope を　新しく増やす (scope によって　表示する画面を変える)
4 JWT に基づいて,再連携 => 保存内容を更新する（spotify/twitter api を 再び叩く）

*/

type Authenticator struct {
	SocialService domain.SocialAccount 
}


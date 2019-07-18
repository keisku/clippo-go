# プロダクト "clippo"

#### 「検索したページをメモと一緒に残しておきたい」

エラー解決や学習のために読んだWebページをページ内容の要約と共に保存しておけるエンジニアのためのサービスです。

# 機能一覧

- ログイン
- 投稿
- ユーザー登録
- 投稿検索

# 使用技術

- Go
- マイクロサービス構成
- gRPC通信
- JWT認証
- Webスクレイピング

# 使用予定の技術

- AWS
  - ECS, ECR, RDS, CloudFormation
- Docker

面接までにサーバへアプリケーションをデプロイしてお見せしたかったのですが、現状ローカル環境で動作するレベルで止まっています。
予定では、上記のインフラでデプロイを考えておりました。

# ブランチに関しての補足

- master
    ブランチではECSでのデプロイを意識した構成です。
- Local
    デモ用にローカル環境でWebアプリケーションが作動する構成です。

# デモGIF（ローカル環境での動作をキャプチャ）

## ユーザー登録〜ログイン（フルスクラッチで開発）

使用している技術
- JWT認証

![ユーザー登録〜ログイン](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-signup-login.gif)

## Webページのクリップ

使用している技術
- Webスクレイピング

![Webページのクリップ](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-clip.gif)

## クリップした記事のタイトル検索

![タイトル検索](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-search.gif)

## クリップした記事のタグ検索

![タグ検索](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-search-tag.gif)

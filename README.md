# clippo
エラー解決や学習のために読んだWebページをページ内容の要約と共に保存しておけるエンジニアのためのサービスです。
*前提として、Webアプリケーションを開発できるという技術を見ていただきたいので、デザイン、サービスとしての質は最低限しか考慮していません。

# 使用技術
- Go
- マイクロサービス構成
- gRPC通信
- JWT認証
- Webスクレイピング

# 使用予定の技術
- AWS
  - ECS, ECR, RDS

サービスをデプロイするためにインフラでは、AWSを使う予定です。

# デモGIF

## ユーザー登録〜ログイン(フルスクラッチで開発)

![ユーザー登録〜ログイン](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-signup-login.gif)

## Webページのクリップ

![Webページのクリップ](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-clip.gif)

## クリップした記事のタイトル検索

![タイトル検索](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-search.gif)

## クリップした記事のタグ検索

![タグ検索](https://github.com/kskumgk63/clippo-go/blob/Local/GIF/clippo-search-tag.gif)

# プロダクト名 "clippo"

#### 「検索したページをメモと一緒に残しておきたい」

エラー解決やプログラミング学習のために読んだWebページをページ内容のメモと共に保存しておけるサービスです。
「Pocket」「はてなブックマーク」を意識しています。ローカル環境では動作確認済みです。

# アーキテクチャ
#### entity
モデル定義

#### repository
サービスからリクエストを受けてDB操作を行う

#### service
HTTPパスに応じたサービスを実行する

#### --pb
Protocol Buffersの定義

# 機能一覧

- ログイン
- ユーザー登録
- 投稿
- タグ作成
- 投稿検索（タイトルまたはタグ）

# 使用技術

- Go
- マイクロサービス構成
- gRPC通信
- JWT認証
- Webスクレイピング

# インフラ構築

- AWS
  - ECS, ECR, RDS, ALB,  CloudFormation
- Docker

現状、サービスは本番環境で動作していませんが、インフラでのデプロイを目指しています。

# デモGIF（ローカル環境での動作をキャプチャ）

### ユーザー登録〜ログイン

#### 使用している技術
- JWT認証

#### 処理の流れ
ユーザー登録
- ユーザーがメールアドレスとパスワードと確認パスワードを入力
    - すでに存在するメールアドレスではないかチェック
    - パスワードが確認パスワードと同じかチェック
    - パスワードを暗号化
- DBへ保存

ログイン
- ユーザーがメールアドレスとパスワードとを入力
    - メールアドレスをもとにDBからユーザーのメールアドレスとパスワードを取得
    - パスワードが正しければJWTトークンをキャッシュに格納
    - トークンがなければログイン画面へリダイレクトとする
- トップ画面を表示

![ユーザー登録〜ログイン](https://github.com/kskumgk63/clippo-go/blob/local/GIF/signup-login-top.gif)

### Webページのクリップ

#### 使用している技術
- Webスクレイピング

#### 処理の流れ
記事のクリップ
- クリップしたい記事のURLをタイプして、「Clip」ボタンを押下
    - URL先のタイトル、イメージ画像、ディスクリプションをスクレイピング
    - 投稿確認画面にて修正がある場合は修正
    - 投稿に対してタグを作成
- DBへ保存

![Webページのクリップ](https://github.com/kskumgk63/clippo-go/blob/master/GIF/createPost.gif)

### クリップした記事の検索

#### 処理の流れ
- ユーザーが入力したタイトルからクリップした記事を検索

![タイトル検索](https://github.com/kskumgk63/clippo-go/blob/master/GIF/searchPostsByTitle.gif)

- ユーザーが入力したタグからクリップした記事を検索

![タイトル検索](https://github.com/kskumgk63/clippo-go/blob/master/GIF/searchPostsByTag.gif)

# go-clean-architecture

ディレクトリは[golang-standards](https://github.com/golang-standards/project-layout/blob/master/README_ja.md)を参照。

## Refs
- [Go言語とClean ArchitectureでAPIサーバを構築する](https://qiita.com/arkuchy/items/874656b33d2e5acdf281#%E3%81%BE%E3%81%A8%E3%82%81)
- [【Go】厳密なClean Architectureとその妥協案](https://qiita.com/arkuchy/items/659a11767912c2ec266d)
- [【Go】基本文法総まとめ](https://qiita.com/k-penguin-sato/items/deaeab18aa416496e273)
- [【Go】テストまで書いて理解するレイヤードアーキテクチャ](https://zenn.dev/7oh/articles/6338a8ccd470c7)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

## Local development
```
make build
make up
```

## Local server
http://localhost:8080/api/v1/tasks

## Clean Architecture
- [実装クリーンアーキテクチャ](https://qiita.com/nrslib/items/a5f902c4defc83bd46b8#application-business-rules)メモ
### Enterprise Business Rules
- ビジネスロジック
- どんなデータを扱うのか
- ただの箱だけではなくロジックもここ

### Application Business Rules
- ソフトウェアは何ができるのか
  - create user
  - read users
  - etc...

### Interface Adapters
- 入力
  - Application Business Rulesに伝えるためのデータ加工
- 永続化
  - データの保存
  - SQL叩く
- 表示
  - 結果の表示

### Frameworks & Drivers
- DBのconnectionの設定
- ルーティング

## まだわかってないこと
- バリデーションはどこがやるのか
  - ユーザー登録で電話番号に文字列が入ってきた
  - method:GET /users/{id}
    - idに存在しないuserのidが入ってきた

- アクセス権限はどこがやるのか
  - ログインしているユーザーだけリクエスト飛ばせる

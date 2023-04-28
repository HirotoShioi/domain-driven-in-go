# GOでドメイン駆動開発

以下の機能を持ったアプリケーションをドメイン駆動開発を用いて開発する

- ユーザーの登録
- ユーザーの削除
- 登録済みユーザーの取得
- ユーザーの更新

ユーザーのエンティティは以下

```
User:
    id: uuid;
    username: string; ユニークキー制約を持たせる
    email: string;
```

# 使うパッケージ

- [sqlc](https://sqlc.dev/)
- [migrate](https://github.com/golang-migrate/migrate)
- [dbdiagram](https://dbdiagram.io/)

## ステップ

- [x] 値オブジェクトを実装する
- [x] エンティティを実装する
- [x] ドメインサービスを実装する
- [x] リポジトリを実装する
- [x] アプリケーションサービスを実装する
- [] migrateを改善
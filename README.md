# go-template by matty

- goのアプリ開発を簡略化するためのものです
- ginのチュートリアルに出てくる、アルバムのCRUDをサンプルとして実装しています

# How to start

1. template repository から新しいリポジトリを作成する
2. ルートディレクトリ配下の`go-template`の文字列を任意の文字列に変える
3. devcontainerで立ち上げる
4. `cd bd/ && make migrate`
5. 開発スタート

# 参考資料
- CloudRunについて
  - https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service?hl=ja

#　やりたいこと
- [ ] インフラのTerraform化
- [ ] 単体テスト作成
  - [ ] 自動テスト
- [ ] マイグレーションをCIに組み込む

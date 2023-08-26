# Bread_Cli

先に
https://github.com/bskcorona-github/pan
こちらのREADMEに記述している事を終わらせておく必要がある。

1.go run main.goをターミナルで叩く
2.ブラウザでhttp://localhost:8080/graphqlにアクセス
3.一覧取得したい場合、
{
  entries {
    id
    name
    createdAt
  }
}

を入力して「▶」ボタンを押下

4.個別取得したい場合
{
  entry(id: "取得したいパンのIDを入力") {
    id
    name
    createdAt
  }
}

を入力して「▶」ボタンを押下

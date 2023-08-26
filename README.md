# Bread_Cli

このリポジトリは、Breadプロジェクトと対話するためのコマンドラインインターフェース（CLI）です。続行する前に、[panリポジトリのREADME](https://github.com/bskcorona-github/pan)で説明されているセットアップを完了させてください。

## 使用方法

1. ターミナルで次のコマンドを実行してください：
   ```
   go run main.go
   ```

2. ウェブブラウザを開き、[http://localhost:8080/graphql](http://localhost:8080/graphql)にアクセスしてください。

3. エントリのリストを取得するには、次のGraphQLクエリを入力し、「▶」ボタンをクリックしてください：
   ```graphql
   {
     entries {
       id
       name
       createdAt
     }
   }
   ```

4. 個々のエントリを取得するには、次のGraphQLクエリを入力し、対象のパンのIDを入力してから「▶」ボタンをクリックしてください：
   ```graphql
   {
     entry(id: "ここに取得したいパンのIDを入力") {
       id
       name
       createdAt
     }
   }
   ```

必ず、「ここに取得したいパンのIDを入力」を実際のパンのIDに置き換えてください。

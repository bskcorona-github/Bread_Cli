# Bread_Cli

このリポジトリは、Breadプロジェクトと対話するためのコマンドラインインターフェース（CLI）です。続行する前に、[panリポジトリのREADME](https://github.com/bskcorona-github/pan)で説明されているセットアップを完了させてください。

## 使用方法
1.  以下のコード行で、PostgreSQL データベースへの接続情報を設定します。これを、ご自身の PostgreSQL の接続情報に合わせて変更してください。
   ```go
   connStr := "user=postgres dbname=postgres sslmode=disable password=tkz2001r"
   ```

2. ターミナルで次のコマンドを実行してください：
   ```
   go run main.go
   ```

3. ウェブブラウザを開き、[http://localhost:8080/graphql](http://localhost:8080/graphql)にアクセスしてください。

4. エントリのリストを取得するには、次のGraphQLクエリを入力し、「▶」ボタンをクリックしてください：
   ```graphql
   {
     entries {
       id
       name
       createdAt
     }
   }
   ```

5. 個々のエントリを取得するには、次のGraphQLクエリを入力し、対象のパンのIDを入力してから「▶」ボタンをクリックしてください：
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

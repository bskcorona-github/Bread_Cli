package graphql

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bskcorona-github/Bread_Cli/internal/database"
	"github.com/bskcorona-github/Bread_Cli/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Server struct {
	Schema *graphql.Schema
	DB     *database.DB
}

// Entry タイプの定義

// entryType の定義
var entryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Entry",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					entry, _ := p.Source.(*model.Entry)
					return entry.Sys.ID, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					entry, _ := p.Source.(*model.Entry)
					return entry.Fields.Name, nil
				},
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					entry, _ := p.Source.(*model.Entry)
					return entry.Sys.CreatedAt, nil
				},
			},
		},
	},
)

func NewServer(db *database.DB) *Server {
	// クエリフィールドを定義
	var entryQuery = &graphql.Field{
		Type: entryType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// クエリの解決ロジックを実装
			id, ok := p.Args["id"].(string)
			if !ok {
				return nil, fmt.Errorf("ID argument is required")
			}

			// データベースから指定されたIDのエントリーを取得する
			entry := &model.Entry{}
			err := db.QueryRow("SELECT id, name, created_at FROM entries WHERE id = $1", id).
				Scan(&entry.Sys.ID, &entry.Fields.Name, &entry.Sys.CreatedAt)
			if err != nil {
				log.Println("Error retrieving entry:", err)
				return nil, err
			}

			return entry, nil
		},
	}

	// rootQuery の定義に entryQuery を追加
	var rootQuery = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"entries": &graphql.Field{
					Type: graphql.NewList(entryType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// データベースからエントリー一覧を取得するロジックを実装...
						fmt.Println("Fetching entries from the database...")
						rows, err := db.Query("SELECT id, name, created_at FROM entries")
						if err != nil {
							log.Println("Error querying entries:", err)
							return nil, err
						}
						defer rows.Close()

						var entries []*model.Entry
						for rows.Next() {
							var entry model.Entry
							err := rows.Scan(&entry.Sys.ID, &entry.Fields.Name, &entry.Sys.CreatedAt)
							if err != nil {
								log.Println("Error scanning entry row:", err)
								return nil, err
							}
							entries = append(entries, &entry)
						}
						fmt.Printf("Fetched %d entries from the database.\n", len(entries))
						return entries, nil
					},
				},
				"entry": entryQuery,
			},
		},
	)

	// GraphQLスキーマをここで定義
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		},
	)
	return &Server{
		Schema: &schema,
		DB:     db,
	}
}

func (s *Server) SetupRoutes() {
	h := handler.New(&handler.Config{
		Schema:   s.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
}

func (s *Server) StartServer() error {
	return http.ListenAndServe(":8080", nil)
}

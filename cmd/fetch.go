package cmd

import (
	"fmt"
	"log"

	"github.com/bskcorona-github/Bread_Cli/internal/database"
	"github.com/bskcorona-github/Bread_Cli/internal/graphql"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch and store bread information from Contentful",
	RunE: func(cmd *cobra.Command, args []string) error {
		accessToken := ""
		spaceID := "2vskphwbz4oc"
		entryIDs := []string{"6QRk7gQYmOyJ1eMG9H4jbB", "41RUO5w4oIpNuwaqHuSwEc", "4Li6w5uVbJNVXYVxWjWVoZ"}

		client := graphql.NewClient(accessToken, spaceID)
		entries, err := client.FetchBreadEntries(entryIDs)
		if err != nil {
			log.Printf("Error fetching bread entries: %v", err)
			return err
		}

		db := database.NewDB()
		defer db.Close()

		for _, entry := range entries {
			fmt.Printf("Processing entry: %s\n", entry.ID) // ログ出力
			dbEntry := database.Bread{
				ID:        entry.ID,
				Name:      entry.Name,
				CreatedAt: entry.CreatedAt,
			}
			if err := db.CreateBread(&dbEntry); err != nil {
				log.Printf("Error creating bread entry: %v", err) // エラーログを出力
				return err
			}
			fmt.Printf("Bread created successfully: %s\n", dbEntry.ID) // ログ出力
		}

		fmt.Println("Data fetched and stored successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

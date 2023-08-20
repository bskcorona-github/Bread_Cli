package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bskcorona-github/Bread_Cli/internal/database"
)

// ログ出力関数
func logInfo(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func logError(err error) {
	log.Printf("[ERROR] %v", err)
}

type Client struct {
	AccessToken string
	SpaceID     string
}

func NewClient(accessToken, spaceID string) *Client {
	return &Client{
		AccessToken: accessToken,
		SpaceID:     spaceID,
	}
}

func (c *Client) FetchBreadEntries(entryIDs []string) ([]database.Bread, error) {
	var entries []database.Bread

	for _, entryID := range entryIDs {
		query := fmt.Sprintf(`
			query {
				entry(id: "%s") {
					id
					name
					createdAt
				}
			}
		`, entryID)
		logInfo("Query: %s", query)

		requestBody := struct {
			Query string `json:"query"`
		}{
			Query: query,
		}

		requestBodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			logError(err) // エラーログを出力
			return nil, err
		}

		req, err := http.NewRequest("POST", "https://graphql.contentful.com/content/v1/spaces/"+c.SpaceID+"/environments/master", bytes.NewBuffer(requestBodyBytes))
		if err != nil {
			logError(err) // エラーログを出力
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logError(err) // エラーログを出力
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logError(err) // エラーログを出力
			return nil, err
		}

		var responseData struct {
			Data struct {
				Entry database.Bread `json:"entry"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &responseData); err != nil {
			logError(err) // エラーログを出力
			return nil, err
		}

		// responseData.Data.Entry を entries に追加
		entries = append(entries, responseData.Data.Entry)
	}

	logInfo("Fetched bread entries: %v", entries) // ログ出力
	return entries, nil
}

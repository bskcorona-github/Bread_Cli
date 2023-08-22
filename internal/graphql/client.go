package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

// type graphqlResponse struct {
// 	Data struct {
// 		Entry struct {
// 			ID        string    `json:"id"`
// 			Name      string    `json:"name"`
// 			CreatedAt time.Time `json:"createdAt"`
// 		} `json:"entry"`
// 	} `json:"data"`
// }

const timeLayout = "2006-01-02T15:04:05.999Z07:00"

func (c *Client) FetchBreadEntries(entryIDs []string) ([]database.Bread, error) {
	var entries []database.Bread

	for _, entryID := range entryIDs {
		query := fmt.Sprintf(`
		query {
			story(id: "%s") {
				sys{
					id
					createdAt
				}
				fields {
					name
				}
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
			logError(err)
			return nil, err
		}

		req, err := http.NewRequest("POST", "https://graphql.contentful.com/content/v1/spaces/"+c.SpaceID, bytes.NewBuffer(requestBodyBytes))
		if err != nil {
			logError(err)
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logError(err)
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logError(err)
			return nil, err
		}
		log.Printf("Response body: %s", body)

		var response struct {
			Data struct {
				Story struct {
					Sys struct {
						ID        string `json:"id"`
						CreatedAt string `json:"createdAt"`
					} `json:"sys"`
					Name string `json:"name"`
				} `json:"story"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &response); err != nil {
			logError(err)
			return nil, err
		}
		log.Printf("ID: %s", response.Data.Story.Sys.ID)
		log.Printf("CreatedAt: %s", response.Data.Story.Sys.CreatedAt)
		log.Printf("Name: %s", response.Data.Story.Name)

		createdAt, err := time.Parse(timeLayout, response.Data.Story.Sys.CreatedAt)
		if err != nil {
			logError(err)
			return nil, err
		}

		entries = append(entries, database.Bread{
			ID:        response.Data.Story.Sys.ID,
			Name:      response.Data.Story.Name,
			CreatedAt: createdAt,
		})
	}
	logInfo("Fetched bread entries: %v", entries)
	return entries, nil
}

// func (c *Client) sendGraphQLRequest(query string) (*graphqlResponse, error) {
// 	requestBody := struct {
// 		Query string `json:"query"`
// 	}{
// 		Query: query,
// 	}

// 	requestBodyBytes, err := json.Marshal(requestBody)
// 	if err != nil {
// 		logError(err)
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", "https://graphql.contentful.com/content/v1/spaces/"+c.SpaceID, bytes.NewBuffer(requestBodyBytes))
// 	if err != nil {
// 		logError(err)
// 		return nil, err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		logError(err)
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logError(err)
// 		return nil, err
// 	}

// 	var responseData graphqlResponse

// 	if err := json.Unmarshal(body, &responseData); err != nil {
// 		logError(err)
// 		return nil, err
// 	}

// 	return &responseData, nil
// }

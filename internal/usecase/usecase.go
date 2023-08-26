package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bskcorona-github/Bread_Cli/internal/database"
	"github.com/bskcorona-github/Bread_Cli/model"
)

type EntryUseCase struct {
	DB *database.DB
}

const (
	apiURL      = "https://cdn.contentful.com/spaces/2vskphwbz4oc/entries"
	accessToken = "WBSUkhUPSIYTnETcosrDj3dpWeqaidhogcsJkNXRL3Y"
)

func (uc *EntryUseCase) checkEntryExists(entryID string) (bool, error) {
	var count int
	err := uc.DB.QueryRow("SELECT COUNT(*) FROM entries WHERE id = $1", entryID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}

func (uc *EntryUseCase) getEntry(entryID string) (*model.Entry, error) {
	client := &http.Client{}

	url := fmt.Sprintf("%s/%s?access_token=%s", apiURL, entryID, accessToken)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entry model.Entry
	err = json.NewDecoder(resp.Body).Decode(&entry)
	if err != nil {
		return nil, err
	}
	fmt.Println("entry.Sys:", entry.Sys)
	fmt.Println("entry.Fields:", entry.Fields)

	// "createdAt" フィールドのパース
	createdAtStr := entry.Sys.CreatedAt
	createdAt, err := time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, err
	}
	entry.Sys.CreatedAt = createdAt.Format("2006-01-02 15:04:05-07:00")

	return &entry, nil
}

func (uc *EntryUseCase) saveEntry(entry *model.Entry) error {
	_, err := uc.DB.Exec("INSERT INTO entries (id, name, created_at) VALUES ($1, $2, $3)", entry.Sys.ID, entry.Fields.Name, entry.Sys.CreatedAt)
	return err
}

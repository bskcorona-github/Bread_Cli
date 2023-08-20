package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

// ログ出力関数
func logInfo(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func logError(err error) {
	log.Printf("[ERROR] %v", err)
}

func NewDB() *DB {
	// データベース接続の設定など
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=tkz2001r sslmode=disable")
	if err != nil {
		panic("Failed to connect to the database.")
	}

	// モデルに基づいたテーブルの自動生成
	if err := db.AutoMigrate(&Bread{}).AddUniqueIndex("idx_unique_id", "id").Error; err != nil {
		logError(err) // エラーログを出力
		panic("Failed to migrate the database.")
	}

	logInfo("Database connection established.") // ログ出力
	return &DB{db}
}

// CreateBread はパン情報をデータベースに保存するメソッドです
func (db *DB) CreateBread(bread *Bread) error {
	if err := db.Create(bread).Error; err != nil {
		logError(err) // エラーログを出力
		return err
	}
	logInfo("Bread created successfully: %s", bread.ID) // ログ出力
	return nil
}

// internal/database/models.go
package database

import "time"

type Bread struct {
	ID        string    `gorm:"not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
}

package migrate

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Perk table structure
type Perk struct {
	// Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt mysql.NullTime

	PerkID      uint `sql:"index"`
	TagLine     string
	Description string `sql:"type:text"`
}

func init() {
	filename := "000001_create_perks"
	migrationList = append(migrationList, &Migration{
		Name: filename,
		Function: func(db *gorm.DB) error {
			return db.CreateTable(&Perk{}).Error
		},
	})
}

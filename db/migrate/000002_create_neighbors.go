package migrate

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Neighbor table structure
type Neighbor struct {
	// Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt mysql.NullTime

	UserID      uint `sql:"index"`
	Neighbor1ID uint
	Neighbor2ID uint
	Neighbor3ID uint
	Neighbor4ID uint
	Neighbor5ID uint
}

func init() {
	filename := "000002_create_neighbors"
	migrationList = append(migrationList, &Migration{
		Name: filename,
		Function: func(db *gorm.DB) error {
			return db.CreateTable(&Neighbor{}).Error
		},
	})
}

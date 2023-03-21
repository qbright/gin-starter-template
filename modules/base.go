package modules

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// This functions are called before creating Base
func (base *Base) BeforeCreate(scope *gorm.DB) error {
	base.ID = uuid.NewV4().String()
	return nil
}

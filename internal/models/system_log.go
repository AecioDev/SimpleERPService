package models

import (
	"time"

	"gorm.io/gorm"
)

// SystemLog representa um log do sistema
type SystemLog struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `json:"user_id"`
	Action     string         `gorm:"size:100;not null" json:"action"`
	EntityType string         `gorm:"size:50" json:"entity_type"`
	EntityID   string         `json:"entity_id"`
	Details    map[string]interface{} `gorm:"type:jsonb" json:"details"`
	IPAddress  string         `gorm:"size:45" json:"ip_address"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

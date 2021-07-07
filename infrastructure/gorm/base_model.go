package gorm

import "time"

type BaseModel struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updated_at"`
}

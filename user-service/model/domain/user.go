package domain

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Name      string     `gorm:"not null" json:"name"`
	Password  string     `gorm:"not null" json:"-"`
	AddressID string     `gorm:"type:decimal(10,2);not null;default:null" json:"address_id"`
	CreatedAt time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "users" // Ensure it matches your actual database table name
}

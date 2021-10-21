package entity

import "time"

type Food struct {
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64     `gorm:"size:100;not null;" json:"user_id"`
	Title       string     `gorm:"size:100;not null;unique" json:"title"`
	Description string     `gorm:"text;not null;" json:"description"`
	FoodImage   string     `gorm:"size:255;null;" json:"food_image"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

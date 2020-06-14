package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type ServiceClient struct {
	gorm.Model
	ClientID string
	ClientSecret string
	Status sql.NullBool
}

type ActivityLog struct {
	gorm.Model
	ServiceClientID int8
	Event string
	Data string `gorm:"text"`
	CreatorID string
	CreatorType string
}
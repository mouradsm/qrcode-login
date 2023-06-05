package models

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

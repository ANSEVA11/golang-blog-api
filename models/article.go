//------ In The Name Of God

package models

import "gorm.io/gorm"

type Article struct{
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Body string `json:"body" gorm:"type:text; not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User User `json:"user" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
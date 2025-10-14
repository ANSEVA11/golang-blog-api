//------ In The Name Of God

package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Articles []Article `json:"articles" gorm:"foreignKey:UserID"`
}

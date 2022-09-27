package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200); not null"`
	Email    string `gorm:"type:varchar(200); not null"`
	Password string `gorm:"type:varchar(200); not null"`
	Products []Product
	UserRole UserRole
}

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(200); not null "`
	Description string  `gorm:"type:varchar(200); not null"`
	Price       float64 `gorm:"type:float(200); not null"`
	UserID      int
}

type Role struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200); not null "`
	UserRole UserRole
}
type UserRole struct {
	gorm.Model
	UserID int
	RoleID int
}

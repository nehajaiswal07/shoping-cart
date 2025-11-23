package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Token    string
}

type Item struct {
	gorm.Model
	Name  string
	Price float64
}

type Cart struct {
	gorm.Model
	UserID uint
	ItemID uint
}

type Order struct {
	gorm.Model
	UserID uint
	CartID uint
}

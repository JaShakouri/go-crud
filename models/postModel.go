package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `validate:"required,min=4,max=64" json:"Title"`
	Body  string `validate:"required,min=64,max=256" json:"Body"`
}

type PostRequest struct {
	Title string `validate:"required,min=4,max=64," json:"Title"`
	Body  string `validate:"required,min=64,max=256," json:"Body"`
}

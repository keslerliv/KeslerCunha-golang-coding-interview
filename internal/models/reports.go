package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Num         string `json: "num"`
	Header      string `json: "header"`
	Description string `json: "description"`
	Terms       string `json: "terms"`
}

package models

import(
	"time"
)

type Visit struct {
	ID uint `gorm:"primaryKey"`
	IP string
	Country string
	City string
	WebsiteTag string
	Timestamp time.Time
}

type VisitInput struct {
	IP string `json:"ip" binding:"required"`
	WebsiteTag string `json:"website_tag" bindng:"required"`
}
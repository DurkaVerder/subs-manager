package models

import "time"

type ServiceSubscription struct {
	Name      string
	Price     int64
	UserID    string
	StartDate time.Time
	EndDate   time.Time
}

type ServiceSubscriptionRequest struct {
	Name      string `json:"service_name"`
	Price     int64  `json:"price"`
	UserID    string `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

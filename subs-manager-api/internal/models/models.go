package models

import "time"

// ServiceSubscription represents a subscription to a service.
type ServiceSubscription struct {
	ID        int64
	Name      string
	Price     int64
	UserID    string
	StartDate time.Time
	EndDate   time.Time
}

// ServiceSubscriptionRequest represents the request structure for creating a service subscription.
type ServiceSubscriptionRequest struct {
	Name      string `json:"service_name"`
	Price     int64  `json:"price"`
	UserID    string `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// ServiceSubscriptionResponse represents the response structure for a service subscription.
type ServiceSubscriptionResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"service_name"`
	Price     int64  `json:"price"`
	UserID    string `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// DataFilter includes filters for subscriptions.
type DataFilter struct {
	UserID      string
	ServiceName string
	StartDate   string
	EndDate     string
}

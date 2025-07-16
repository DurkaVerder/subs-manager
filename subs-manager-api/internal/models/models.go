package models

import (
	"time"
)

// ServiceSubscription represents a subscription to a service.
type ServiceSubscription struct {
	ID        int64
	Name      string
	Price     int64
	UserID    string
	StartDate time.Time
	EndDate   time.Time
}

func (s *ServiceSubscription) ToResponse() ServiceSubscriptionResponse {
	return ServiceSubscriptionResponse{
		ID:        s.ID,
		Name:      s.Name,
		Price:     s.Price,
		UserID:    s.UserID,
		StartDate: s.StartDate.Format("01-2006"),
		EndDate:   s.EndDate.Format("01-2006"),
	}
}

// ServiceSubscriptionRequest represents the request structure for creating a service subscription.
type ServiceSubscriptionRequest struct {
	Name      string `json:"service_name"`
	Price     int64  `json:"price"`
	UserID    string `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (r *ServiceSubscriptionRequest) ToModel() (ServiceSubscription, error) {
	startDate, err := time.Parse("01-2006", r.StartDate)
	if err != nil {
		return ServiceSubscription{}, err
	}
	endDate, err := time.Parse("01-2006", r.EndDate)
	if err != nil {
		return ServiceSubscription{}, err
	}

	return ServiceSubscription{
		Name:      r.Name,
		Price:     r.Price,
		UserID:    r.UserID,
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
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
	StartDate   time.Time
	EndDate     time.Time
}

// TotalPriceResponse represents the response structure for the total price of subscriptions.
type TotalPriceResponse struct {
	TotalPrice int64 `json:"total_price"`
}

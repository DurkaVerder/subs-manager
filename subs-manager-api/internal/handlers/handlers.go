package handlers

import "subs-manager-api/internal/models"

type SubscribeService interface {
	GetSubscriptionFilter(data models.DataFilter) ([]models.ServiceSubscriptionResponse, error)
	GetSubscription(serviceName, userID string) (models.ServiceSubscriptionResponse, error)
	GetSubscriptions(userID string) ([]models.ServiceSubscriptionResponse, error)
	GetTotalPriceByFilters(data models.DataFilter) (models.TotalPriceResponse, error)
	CreateSubscription(subscription models.ServiceSubscriptionRequest) error
	UpdateSubscription(subscription models.ServiceSubscriptionRequest) error
	DeleteSubscription(serviceName, userID string) error
}

type Handler struct {
	subscribeService SubscribeService
}

func NewHandler(subscribeService SubscribeService) *Handler {
	return &Handler{subscribeService: subscribeService}
}

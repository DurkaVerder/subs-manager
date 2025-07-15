package handlers

import "subs-manager-api/internal/models"

type SubscribeService interface {
	GetSubscriptionFilter(data models.DataFilter) ([]models.ServiceSubscription, error)
	GetSubscription(serviceName, userID string) (models.ServiceSubscription, error)
	GetSubscriptions(userID string) ([]models.ServiceSubscription, error)
	CreateSubscription(subscription models.ServiceSubscription) error
	UpdateSubscription(subscription models.ServiceSubscription) error
	DeleteSubscription(serviceName, userID string) error
}

type Handler struct {
	subscribeService SubscribeService
}

func NewHandler(subscribeService SubscribeService) *Handler {
	return &Handler{subscribeService: subscribeService}
}

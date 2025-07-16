package subscribe

import (
	"strconv"
	"subs-manager-api/internal/models"
)

type Storage interface {
	GetSubscription(serviceName, userID string) (models.ServiceSubscription, error)
	GetSubscriptions(userID string) ([]models.ServiceSubscription, error)
	GetSubscriptionFilter(filters models.DataFilter) ([]models.ServiceSubscription, error)
	CreateSubscription(subscription models.ServiceSubscription) error
	UpdateSubscription(subscription models.ServiceSubscription) error
	DeleteSubscription(ID int64) error
}

type SubscribeService struct {
	storage Storage
}

func NewSubscribeService(storage Storage) *SubscribeService {
	return &SubscribeService{storage: storage}
}

func (s *SubscribeService) GetSubscription(serviceName, userID string) (models.ServiceSubscriptionResponse, error) {

	subscribe, err := s.storage.GetSubscription(serviceName, userID)
	if err != nil {
		return models.ServiceSubscriptionResponse{}, nil
	}

	return subscribe.ToResponse(), nil

}

func (s *SubscribeService) GetSubscriptions(userID string) ([]models.ServiceSubscriptionResponse, error) {
	subscribes, err := s.storage.GetSubscriptions(userID)
	if err != nil {
		return nil, err
	}

	var responses []models.ServiceSubscriptionResponse
	for _, sub := range subscribes {
		responses = append(responses, sub.ToResponse())
	}

	return responses, nil
}

func (s *SubscribeService) CreateSubscription(subscription models.ServiceSubscriptionRequest) error {

	models, err := subscription.ToModel()
	if err != nil {
		return err
	}

	if err := s.storage.CreateSubscription(models); err != nil {
		return err
	}

	return nil
}

func (s *SubscribeService) UpdateSubscription(subscription models.ServiceSubscriptionRequest) error {

	models, err := subscription.ToModel()
	if err != nil {
		return err
	}

	if err := s.storage.UpdateSubscription(models); err != nil {
		return err
	}

	return nil
}

func (s *SubscribeService) DeleteSubscription(ID string) error {
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteSubscription(id); err != nil {
		return err
	}

	return nil
}

func (s *SubscribeService) GetSubscriptionByFilter(data models.DataFilter) ([]models.ServiceSubscriptionResponse, error) {
	subscriptions, err := s.storage.GetSubscriptionFilter(data)
	if err != nil {
		return nil, err
	}

	var responses []models.ServiceSubscriptionResponse
	for _, sub := range subscriptions {
		responses = append(responses, sub.ToResponse())
	}

	return responses, nil
}

func (s *SubscribeService) GetTotalPriceByFilters(data models.DataFilter) (models.TotalPriceResponse, error) {
	subscriptions, err := s.storage.GetSubscriptionFilter(data)
	if err != nil {
		return models.TotalPriceResponse{}, err
	}

	totalPrice := s.getPriceSubscriptions(subscriptions)
	return models.TotalPriceResponse{TotalPrice: totalPrice}, nil
}

func (s *SubscribeService) getPriceSubscriptions(subs []models.ServiceSubscription) int64 {
	var total int64
	for _, subscription := range subs {
		total += subscription.Price
	}
	return total
}

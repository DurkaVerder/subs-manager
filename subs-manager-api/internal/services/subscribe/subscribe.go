package subscribe

import "subs-manager-api/internal/models"

type Storage interface {
	GetSubscription(serviceName, userID string) (models.ServiceSubscription, error)
	GetSubscriptions(userID string) ([]models.ServiceSubscription, error)
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

func (s SubscribeService) GetSubscriptionFilter(data models.DataFilter) ([]models.ServiceSubscription, error) {
	// Coming soon
	return nil, nil
}

func (s SubscribeService) GetSubscription(serviceName, userID string) (models.ServiceSubscription, error) {

	subscribe, err := s.storage.GetSubscription(serviceName, userID)
	if err != nil {
		return models.ServiceSubscription{}, nil
	}

	return subscribe, nil

}

func (s SubscribeService) GetSubscriptions(userID string) ([]models.ServiceSubscription, error) {
	subscribes, err := s.storage.GetSubscriptions(userID)
	if err != nil {
		return nil, err
	}

	return subscribes, nil
}

func (s SubscribeService) CreateSubscription(subscription models.ServiceSubscription) error {
	if err := s.storage.CreateSubscription(subscription); err != nil {
		return err
	}

	return nil
}

func (s SubscribeService) UpdateSubscription(subscription models.ServiceSubscription) error {
	if err := s.storage.UpdateSubscription(subscription); err != nil {
		return err
	}

	return nil
}

func (s SubscribeService) DeleteSubscription(ID int64) error {
	if err := s.storage.DeleteSubscription(ID); err != nil {
		return err
	}

	return nil
}

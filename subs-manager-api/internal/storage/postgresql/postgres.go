package postgresql

import (
	"database/sql"
	"subs-manager-api/internal/models"
)

const (
	GetSubscriptionQuery    = `SELECT * FROM subscriptions WHERE service_name = $1 AND user_id = $2`
	GetSubscriptionsQuery   = `SELECT * FROM subscriptions WHERE user_id = $1`
	CreateSubscriptionQuery = `INSERT INTO subscriptions (service_name, user_id, price, start_date, end_date) VALUES ($1, $2, $3, $4, $5)`
	UpdateSubscriptionQuery = `UPDATE subscriptions SET price = $1, start_date = $2, end_date = $3 WHERE id = $4`
	DeleteSubscriptionQuery = `DELETE FROM subscriptions WHERE id = $1`

	GetSubscriptionFilterQuery = `SELECT * FROM subscriptions WHERE user_id = $1 AND start_date >= $2 AND end_date <= $3 AND (service_name = $4 OR $4 IS NULL)`
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) GetSubscription(serviceName, userID string) (models.ServiceSubscription, error) {
	var subscription models.ServiceSubscription
	err := p.db.QueryRow(GetSubscriptionQuery, serviceName, userID).Scan(&subscription.Name, &subscription.UserID, &subscription.Price, &subscription.StartDate, &subscription.EndDate)
	if err != nil {
		return models.ServiceSubscription{}, err
	}
	return subscription, nil
}

func (p *Postgres) GetSubscriptions(userID string) ([]models.ServiceSubscription, error) {
	rows, err := p.db.Query(GetSubscriptionsQuery, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var subscriptions []models.ServiceSubscription
	for rows.Next() {
		var subscription models.ServiceSubscription
		if err := rows.Scan(&subscription.Name, &subscription.UserID, &subscription.Price, &subscription.StartDate, &subscription.EndDate); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (p *Postgres) CreateSubscription(subscription models.ServiceSubscription) error {
	_, err := p.db.Exec(CreateSubscriptionQuery, subscription.Name, subscription.UserID, subscription.Price, subscription.StartDate, subscription.EndDate)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateSubscription(subscription models.ServiceSubscription) error {
	_, err := p.db.Exec(UpdateSubscriptionQuery, subscription.Price, subscription.StartDate, subscription.EndDate, subscription.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteSubscription(ID int64) error {

	_, err := p.db.Exec(DeleteSubscriptionQuery, ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetSubscriptionFilter(filters models.DataFilter) ([]models.ServiceSubscription, error) {
	rows, err := p.db.Query(GetSubscriptionFilterQuery, filters.UserID, filters.StartDate, filters.EndDate, filters.ServiceName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribes []models.ServiceSubscription
	for rows.Next() {
		var subscription models.ServiceSubscription
		if err := rows.Scan(&subscription.Name, &subscription.UserID, &subscription.Price, &subscription.StartDate, &subscription.EndDate); err != nil {
			return nil, err

		}
		subscribes = append(subscribes, subscription)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscribes, nil
}

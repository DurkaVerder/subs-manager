

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    price INTEGER NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL CHECK (end_date > start_date),
    UNIQUE (user_id, service_name)
);

CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
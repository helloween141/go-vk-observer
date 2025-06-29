-- +goose Up
-- +goose StatementBegin
CREATE TABLE telegram_notifications (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL,
    entity_id INT NOT NULL,
    last_post_date BIGINT,
    checked_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE telegram_notifications;
-- +goose StatementEnd

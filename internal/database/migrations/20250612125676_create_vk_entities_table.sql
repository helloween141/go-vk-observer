-- +goose Up
-- +goose StatementBegin
CREATE TYPE entity_type AS ENUM('WALL', 'TOPIC');

CREATE TABLE vk_entities (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    type ENTITY_TYPE NOT NULL,
    name VARCHAR(500) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vk_entities;
DROP TYPE entity_type;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE vk_entities DROP COLUMN type;
DROP TYPE ENTITY_TYPE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TYPE entity_type AS ENUM('WALL', 'TOPIC');
ALTER TABLE vk_entities ADD type ENTITY_TYPE
-- +goose StatementEnd

-- name: CreateVkEntity :one
INSERT INTO vk_entities (slug, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetVkEntityBySlug :one
SELECT * FROM vk_entities WHERE slug=$1;
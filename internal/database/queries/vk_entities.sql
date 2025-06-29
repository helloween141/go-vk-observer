-- name: CreateVkEntity :one
INSERT INTO vk_entities (slug, name, type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetVkEntityBySlug :one
SELECT * FROM vk_entities WHERE slug=$1;

-- name: IsVkEntityExists :one
SELECT EXISTS(SELECT 1 FROM vk_entities WHERE id=$1);

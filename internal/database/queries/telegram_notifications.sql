-- name: GetTelegramNotificationList :many
SELECT * FROM telegram_notifications
JOIN vk_entities ve on telegram_notifications.entity_id = ve.id
WHERE telegram_notifications.checked_at IS NULL OR
      telegram_notifications.checked_at < now() - INTERVAL '1 hour';

-- name: CreateTelegramNotification :one
INSERT INTO telegram_notifications (telegram_id, entity_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteTelegramNotification :one
DELETE FROM telegram_notifications WHERE telegram_id=$1 AND entity_id=$2
RETURNING true;

-- name: UpdateTelegramNotification :one
UPDATE telegram_notifications
SET checked_at=now(), last_post_date=$1
WHERE telegram_id=$2 AND entity_id=$3
RETURNING *;

-- name: GetTelegramNotificationsByTelegramID :many
SELECT * FROM telegram_notifications
JOIN vk_entities ve on telegram_notifications.entity_id = ve.id
WHERE telegram_notifications.telegram_id=$1;

-- name: IsTelegramNotificationExists :one
SELECT EXISTS(SELECT 1 FROM telegram_notifications WHERE telegram_id=$1 AND entity_id=$2);

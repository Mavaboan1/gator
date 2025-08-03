-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
       )
RETURNING *;

-- name: GetUserFeedName :one
SELECT *
FROM users WHERE id = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: DeleteFeed :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;
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

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1;

-- name: GetFeedFromURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUsernames :many
SELECT f.*, u.name AS username FROM feeds AS f JOIN users AS u ON f.user_id = u.ID;

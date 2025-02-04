-- name: CreateFeedFollow :one
WITH new_feed AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT new_feed.*, feeds.name AS feed_name, users.name AS user_name
FROM new_feed
INNER JOIN feeds ON feeds.id = new_feed.feed_id
INNER JOIN users ON users.id = new_feed.user_id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;

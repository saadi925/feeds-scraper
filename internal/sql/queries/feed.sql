-- name: CreateFeed :one 
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id,created_at,updated_at,user_id,feed_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id=$1;

-- name: DeleteFeedFollow :exec
SELECT * FROM feed_follows WHERE id=$1 AND user_id=$2 ;


-- name: GetNextFeedsToFetch :many 
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkedFeedAsFetch :one
UPDATE feeds 
SET updated_at=NOW(),
last_fetched_at=NOW()
WHERE user_id=$1 
RETURNING *;
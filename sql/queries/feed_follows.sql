-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (
    id, created_at, updated_at, user_id, feed_id
  ) VALUES (
  $1,$2,$3,$4,$5
) RETURNING *
)
SELECT iff.* , u.name as username, f.name as feedname
FROM inserted_feed_follow AS iff 
INNER JOIN users AS u ON iff.user_id = u.id
INNER JOIN feeds as f ON iff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, u.name AS username, f.name AS feedname
FROM feed_follows AS ff
INNER JOIN users as u ON ff.user_id = u.id
INNER JOIN feeds as f ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: Unfollow :one
DELETE 
FROM feed_follows AS ff
WHERE ff.user_id = $1 AND ff.feed_id= $2
RETURNING *;

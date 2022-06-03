-- name: CreateStory :one
INSERT INTO stories (
    id,
    user_id,
    party_id,
    url,
    tagged_friends
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteStory :exec
DELETE FROM stories
WHERE id = $1  AND user_id = $2;

-- name: GetStory :one
SELECT * FROM stories
WHERE id = $1 LIMIT 1;

-- name: GetStoryByUser :many
SELECT * FROM stories
WHERE user_id = $1
ORDER BY id desc
LIMIT $2
OFFSET $3;

-- name: GetStoryByParty :many
SELECT * FROM stories
WHERE party_id = $1
ORDER BY id desc
LIMIT $2
OFFSET $3;

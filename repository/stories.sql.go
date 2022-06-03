// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: stories.sql

package repository

import (
	"context"
)

const createStory = `-- name: CreateStory :one
INSERT INTO stories (
    id,
    user_id,
    party_id,
    url,
    tagged_friends
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, user_id, party_id, url, tagged_friends
`

type CreateStoryParams struct {
	ID            string
	UserID        string
	PartyID       string
	Url           string
	TaggedFriends []string
}

func (q *Queries) CreateStory(ctx context.Context, arg CreateStoryParams) (Story, error) {
	row := q.db.QueryRow(ctx, createStory,
		arg.ID,
		arg.UserID,
		arg.PartyID,
		arg.Url,
		arg.TaggedFriends,
	)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PartyID,
		&i.Url,
		&i.TaggedFriends,
	)
	return i, err
}

const deleteStory = `-- name: DeleteStory :exec
DELETE FROM stories
WHERE id = $1  AND user_id = $2
`

type DeleteStoryParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteStory(ctx context.Context, arg DeleteStoryParams) error {
	_, err := q.db.Exec(ctx, deleteStory, arg.ID, arg.UserID)
	return err
}

const getStory = `-- name: GetStory :one
SELECT id, user_id, party_id, url, tagged_friends FROM stories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStory(ctx context.Context, id string) (Story, error) {
	row := q.db.QueryRow(ctx, getStory, id)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PartyID,
		&i.Url,
		&i.TaggedFriends,
	)
	return i, err
}

const getStoryByParty = `-- name: GetStoryByParty :many
SELECT id, user_id, party_id, url, tagged_friends FROM stories
WHERE party_id = $1
ORDER BY id desc
LIMIT $2
OFFSET $3
`

type GetStoryByPartyParams struct {
	PartyID string
	Limit   int32
	Offset  int32
}

func (q *Queries) GetStoryByParty(ctx context.Context, arg GetStoryByPartyParams) ([]Story, error) {
	rows, err := q.db.Query(ctx, getStoryByParty, arg.PartyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Story
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PartyID,
			&i.Url,
			&i.TaggedFriends,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStoryByUser = `-- name: GetStoryByUser :many
SELECT id, user_id, party_id, url, tagged_friends FROM stories
WHERE user_id = $1
ORDER BY id desc
LIMIT $2
OFFSET $3
`

type GetStoryByUserParams struct {
	UserID string
	Limit  int32
	Offset int32
}

func (q *Queries) GetStoryByUser(ctx context.Context, arg GetStoryByUserParams) ([]Story, error) {
	rows, err := q.db.Query(ctx, getStoryByUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Story
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PartyID,
			&i.Url,
			&i.TaggedFriends,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
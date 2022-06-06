package repository

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	TABLE_NAME = "stories"
)

type StoryRepository struct {
	pool    *pgxpool.Pool
	querier Querier
}

func NewStoryRepository(dbUser, dbPW, dbName, dbHost string, dbPort uint16) (*StoryRepository, error) {
	urlStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPW, dbHost, fmt.Sprint(dbPort), dbName)
	pgURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	connURL := *pgURL
	if connURL.Scheme == "cockroachdb" {
		connURL.Scheme = "postgres"
	}

	c, err := pgxpool.ParseConfig(connURL.String())
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}

	err = validateSchema(connURL)
	if err != nil {
		log.Printf("Schema validation error: %v", err)
	}

	return &StoryRepository{
		pool:    pool,
		querier: New(pool),
	}, nil
}

func (d StoryRepository) Close() {
	d.pool.Close()
}

func (r StoryRepository) CreateStory(ctx context.Context, arg CreateStoryParams) (Story, error) {
	return r.querier.CreateStory(ctx, arg)
}

func (r StoryRepository) DeleteStory(ctx context.Context, arg DeleteStoryParams) error {
	return r.querier.DeleteStory(ctx, arg)
}

func (r StoryRepository) GetStory(ctx context.Context, id string) (Story, error) {
	return r.querier.GetStory(ctx, id)
}

func (r StoryRepository) GetStoryByParty(ctx context.Context, arg GetStoryByPartyParams) ([]Story, error) {
	if arg.Limit == 0 {
		arg.Limit = 10
	}

	return r.querier.GetStoryByParty(ctx, arg)
}

func (r StoryRepository) GetStoryByUser(ctx context.Context, arg GetStoryByUserParams) ([]Story, error) {
	if arg.Limit == 0 {
		arg.Limit = 10
	}

	return r.querier.GetStoryByUser(ctx, arg)
}

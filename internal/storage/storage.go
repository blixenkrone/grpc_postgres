package storage

import (
	"context"
	"fmt"

	"github.com/blixenkrone/lea/internal/storage/postgres"
	"github.com/blixenkrone/lea/internal/storage/postgres/learnings"
)

type LearningsStore struct {
	db      postgres.DB
	querier *learnings.Queries
}

func NewLearningStore(connStr string) (LearningsStore, error) {
	pg, err := postgres.NewStore(connStr)
	if err != nil {
		return LearningsStore{}, fmt.Errorf("error creating postgres store: %w", err)
	}
	if err := pg.RunMigrations(); err != nil {
		return LearningsStore{}, fmt.Errorf("error running migrations: %w", err)
	}

	querier := learnings.New(pg.DB())
	return LearningsStore{pg, querier}, nil
}

func (s LearningsStore) Close() error {
	return s.db.Close()
}

func (s LearningsStore) GetCourses(ctx context.Context) ([]learnings.Course, error) {
	c, err := s.querier.ListCourses(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// func (s LearningsStore) AddCourse(ctx context.Context, l learningsv1.Material) {}

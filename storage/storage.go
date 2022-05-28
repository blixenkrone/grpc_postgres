package storage

import (
	"context"
	"errors"
	"fmt"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/blixenkrone/lea/storage/postgres/learnings"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/uuid"
)

type LearningsStore struct {
	db      postgres.DB
	querier *learnings.Queries
}

func NewLearningStore(db postgres.DB) (LearningsStore, error) {
	if err := db.RunMigrations("./storage/postgres/migrations"); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return LearningsStore{}, fmt.Errorf("error running migrations: %w", err)
		}
	}

	querier := learnings.New(db.DB())
	return LearningsStore{db, querier}, nil
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

func (s LearningsStore) AddCourse(ctx context.Context, l *learningsv1.Course) (learnings.Course, error) {
	p := learnings.AddCourseParams{
		ID:         uuid.New(),
		IsActive:   l.IsActive,
		CourseName: l.Name,
		CreatedAt:  l.CreatedAt.AsTime(),
		UpdatedAt:  l.UpdatedAt.AsTime(),
	}
	return s.querier.AddCourse(ctx, p)
}

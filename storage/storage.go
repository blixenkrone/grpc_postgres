package storage

import (
	"context"
	"fmt"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/blixenkrone/lea/storage/postgres/learnings"
	"github.com/google/uuid"
)

type LearningsStore struct {
	db      postgres.DB
	querier *learnings.Queries
}

type LearningsReadWriter interface {
	GetCourses(ctx context.Context) ([]learnings.Course, error)
}

var _ LearningsReadWriter = LearningsStore{}

func NewLearningStore(db postgres.DB) (LearningsStore, error) {
	if err := db.RunMigrations(); err != nil {
		return LearningsStore{}, fmt.Errorf("error running migrations: %w", err)
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
	}
	c, err := s.querier.AddCourse(ctx, p)
	if err != nil {
		return learnings.Course{}, err
	}
	return c, nil
}

package storage

import (
	"context"
	"errors"
	"fmt"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/learnings/v1"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/blixenkrone/lea/storage/postgres/learnings"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type LearningsStore struct {
	log     logrus.FieldLogger
	pg      postgres.DB
	querier *learnings.Queries
}

func NewLearningStore(l logrus.FieldLogger, db postgres.DB) (LearningsStore, error) {
	querier := learnings.New(db.DB())
	return LearningsStore{l, db, querier}, nil
}

func (s LearningsStore) MigrateUp(srcpath string) error {
	if err := s.pg.RunMigrations(srcpath); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("error running migrations: %w", err)
		}
		s.log.Warnf("ran migrations with no change")
	}
	return nil
}

func (s LearningsStore) Close() error {
	return s.pg.Close()
}

func (s LearningsStore) GetCourses(ctx context.Context) ([]learnings.Course, error) {
	return s.querier.ListCourses(ctx)
}

func (s LearningsStore) AddCourse(ctx context.Context, l *learningsv1.Course) (learnings.Course, error) {
	id, err := uuid.Parse(l.Id)
	if err != nil {
		return learnings.Course{}, err
	}
	p := learnings.AddCourseParams{
		ID:         id,
		IsActive:   l.IsActive,
		CourseName: l.Name,
		CreatedAt:  l.CreatedAt.AsTime(),
		UpdatedAt:  l.UpdatedAt.AsTime(),
	}
	tx, err := s.pg.DB().Begin()
	if err != nil {
		return learnings.Course{}, err
	}
	c, err := s.querier.WithTx(tx).AddCourse(ctx, p)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			s.log.Errorf("error rolling back tx: %v", err)
		}
		return learnings.Course{}, fmt.Errorf("error adding course: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return learnings.Course{}, fmt.Errorf("error commit tx: %w", err)
	}

	return c, nil
}

func (s LearningsStore) GetModule(ctx context.Context, moduleID uuid.UUID) (learnings.Module, error) {
	return s.querier.GetModule(ctx, moduleID)
}

func (s LearningsStore) GetModules(ctx context.Context, moduleID uuid.UUID) (learnings.Module, error) {
	return s.querier.Get
}

package grpc

import (
	"context"
	"testing"

	"github.com/blixenkrone/lea/docker"
	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// ! WIP
func TestAddCourse(t *testing.T) {
	l := logrus.New()
	ctx := context.Background()
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			a := assert.New(t)

			pool, err := docker.NewPool()
			a.NoError(err)
			resource, err := pool.NewPostgres("grpcressource")
			a.NoError(err)
			defer resource.Teardown()

			db := postgres.NewFromConn(resource.Container())
			store, err := storage.NewLearningStore(l, db)
			a.NoError(err)

			err = store.MigrateUp("../storage/postgres/migrations")
			a.NoError(err)

			server := NewServer(l, store)
			course, err := server.AddCourse(ctx, &learningsv1.AddCourseRequest{})
			a.NoError(err)
			l.Infof("course id: %s", course.Id)
		})
	}
}

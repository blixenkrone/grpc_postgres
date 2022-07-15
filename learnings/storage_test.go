package learnings

import (
	"testing"

	"github.com/blixenkrone/lea/docker"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLearningsStoreMigrations(t *testing.T) {
	t.Run("Migrations schema runs successfully", func(t *testing.T) {
		a := assert.New(t)
		l := logrus.New()

		pool, err := docker.NewPool()
		a.NoError(err)
		pg, err := pool.Postgres("testdb")
		a.NoError(err)

		pgdb := postgres.NewFromConn(pg.Container())
		ls, err := NewLearningStore(l, pgdb)
		a.NoError(err)
		err = ls.Close()
		a.NoError(err)
	})
}

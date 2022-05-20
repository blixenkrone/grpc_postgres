package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresClient(t *testing.T) {
	t.Run("Docker postgres ressource", func(t *testing.T) {
		a := assert.New(t)
		p, err := NewPool()
		a.NoError(err)
		r, err := p.NewPostgres("testdb")
		a.NoError(err)
		err = r.Teardown()
		a.NoError(err)
	})
}

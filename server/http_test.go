package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/blixenkrone/lea/docker"
	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage"
	"github.com/blixenkrone/lea/storage/postgres"
)

func TestGetCourse(t *testing.T) {
	a := assert.New(t)
	p, err := docker.NewPool()
	a.NoError(err)

	testCases := []struct {
		desc   string
		wantId string
		seedFn func(t *testing.T, s storage.LearningsStore)
	}{
		{
			desc:   "successfully returns a course from storage",
			wantId: "1",
			seedFn: func(t *testing.T, s storage.LearningsStore) {

			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			a := assert.New(t)
			l := logrus.New()
			rr := httptest.NewRecorder()

			pgr, err := p.NewPostgres("learningsdb")
			a.NoError(err)
			defer pgr.Teardown()
			db := pgr.Container()

			pgdb := postgres.NewFromConn(db)
			store, err := storage.NewLearningStore(pgdb)
			a.NoError(err)

			req := httptest.NewRequest(http.MethodGet, "/course", nil)
			req = mux.SetURLVars(req, map[string]string{"id": "1"})

			s := NewServer(l, ":8080", store)
			handler := s.getCourse()
			handler(rr, req)

			var c learningsv1.Course
			err = json.NewDecoder(rr.Body).Decode(&c)
			a.NoError(err)
			a.Equal(tc.wantId, c.Id, "got %v", &c)
		})
	}
}

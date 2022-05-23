package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger      logrus.FieldLogger
	srv         http.Server
	learningsDB storage.LearningsStore
}

func NewServer(l logrus.FieldLogger, addr string, ldb storage.LearningsStore) Server {
	r := mux.NewRouter()
	srv := http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       time.Second * 20,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 20,
		IdleTimeout:       time.Second * 20,
		MaxHeaderBytes:    1 << 20,
	}

	s := Server{l, srv, ldb}
	s.registerRoutes(r)

	return s
}

// TODO: improve loggerMW with better logging
func (s Server) loggerMW(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Infof("calling %s w method %s", r.URL, r.Method)
		h(rw, r)
	}
}

func (s Server) registerRoutes(fh *mux.Router) {
	routes := map[string]struct {
		fn     http.HandlerFunc
		method string
	}{
		"/ping":        {s.pong(), http.MethodGet},
		"/course/{id}": {s.getCourse(), http.MethodGet},
	}
	for k, v := range routes {
		v.fn = s.loggerMW(v.fn)
		fh.HandleFunc(k, v.fn).Methods(v.method)
	}
}

func (s Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s Server) ShutDown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s Server) pong() http.HandlerFunc {
	return func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte("PONG"))
	}
}

func (s Server) getAuth() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

// POST /module/{id}/material
func (s Server) postMaterial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GET /module/{id}/material
func (s Server) getMaterial() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

// POST /module/{id}/material/
func (s Server) addMaterial() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s Server) addCourse() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s Server) getCourse() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			http.Error(rw, "no course ID provided", http.StatusBadRequest)
			return
		}
		c := learningsv1.Course{
			Id:       id,
			IsActive: true,
			Name:     "hello world!",
			Modules: []*learningsv1.Module{
				{
					Id:       id,
					CourseId: uuid.New().String(),
					Level:    learningsv1.Module_LEVEL_BEGINNER,
					Material: []*learningsv1.Material{},
				},
			},
		}
		if err := json.NewEncoder(rw).Encode(&c); err != nil {
			panic(err)
		}
	}
}

package docker

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Pool struct {
	pool *dockertest.Pool
}

func NewPool() (Pool, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return Pool{}, err
	}
	pool.MaxWait = 120 * time.Second
	return Pool{pool}, nil
}

type Resource[T any] struct {
	r         *dockertest.Resource
	container T
}

func (p Pool) Minio(dbname string) (*Resource[*minio.Client], error) {
	env := []string{
		"MINIO_ACCESS_KEY=access-key",
		"MINIO_SECRET_KEY=secret-key",
	}

	runOpts := dockertest.RunOptions{
		Repository: "minio/minio",
		Tag:        "latest",
		Env:        env,
		Cmd:        []string{"server", "/data"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("9000/tcp"): {{HostPort: "9000"}},
		},
	}
	config := func(cfg *docker.HostConfig) {
		cfg.AutoRemove = true
		cfg.RestartPolicy = docker.RestartPolicy{Name: "no"}
	}
	resource, err := p.pool.RunWithOptions(&runOpts, config)
	if err != nil {
		return nil, fmt.Errorf("error starting pg resource: %w", err)
	}
	resource.Expire(120)

	endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp"))

	initFn := func() error {
		url := fmt.Sprintf("http://%s/minio/health/live", endpoint)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return errors.New("status code not OK")
		}
		return nil
	}

	if err := p.pool.Retry(initFn); err != nil {
		return nil, err
	}

	c, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("access-key", "secret-key", ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &Resource[*minio.Client]{resource, c}, nil
}

func (p Pool) Postgres(dbname string) (*Resource[*sql.DB], error) {
	env := []string{
		"POSTGRES_USER=admin",
		"POSTGRES_PASSWORD=password",
		fmt.Sprintf("POSTGRES_DB=%s", dbname),
		"listen_addresses = '*'",
	}

	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14-alpine",
		Env:        env,
	}
	config := func(cfg *docker.HostConfig) {
		cfg.AutoRemove = true
		cfg.RestartPolicy = docker.RestartPolicy{Name: "no"}
	}
	resource, err := p.pool.RunWithOptions(&runOpts, config)
	if err != nil {
		return nil, fmt.Errorf("error starting pg resource: %w", err)
	}
	// Tell docker to hard kill the container in 120 seconds
	resource.Expire(120)

	var pgdb *sql.DB
	initFn := func() error {
		hostAndPort := resource.GetHostPort("5432/tcp")
		// parts := strings.Split(hostAndPort, ":")
		// connStr := fmt.Sprintf("user=admin password=password host=%s port=%s dbname=%s", parts[0], parts[1], dbname)
		connStr := fmt.Sprintf("postgres://admin:password@%s/%s?sslmode=disable", hostAndPort, dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return fmt.Errorf("error creating postgres store: %w", err)
		}
		pgdb = db
		return db.Ping()
	}

	if err := p.pool.Retry(initFn); err != nil {
		return nil, err
	}

	return &Resource[*sql.DB]{resource, pgdb}, nil
}

func (r Resource[T]) Teardown() error {
	return r.r.Close()
}

func (r Resource[T]) Container() T {
	return r.container
}

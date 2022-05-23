package docker

import (
	"database/sql"
	"fmt"
	"time"

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

// TODO: minio

func (p Pool) NewPostgres(dbname string) (*Resource[*sql.DB], error) {
	env := []string{
		"POSTGRES_USER=admin",
		"POSTGRES_PASSWORD=password",
		fmt.Sprintf("POSTGRES_DB=%s", dbname),
		"listen_addresses = '*'",
	}

	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
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
	storeFn := func() error {
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

	if err := p.pool.Retry(storeFn); err != nil {
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

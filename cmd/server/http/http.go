package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blixenkrone/lea/server/http"
	"github.com/blixenkrone/lea/storage"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	addrPort = flag.String("p", ":8080", "which port for server")
)

// TODO: implement gRPC server and gRPC>REST gateway directly in this binary
// https://blog.logrocket.com/guide-to-grpc-gateway/
// https://github.com/grpc-ecosystem/grpc-gateway
func main() {
	l := logrus.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	pgConnStr, ok := os.LookupEnv("SUPABASE_PG_CONN_STRING")
	if !ok {
		panic("no env found for pg conn str")
	}

	db, err := postgres.NewFromConnectionString(pgConnStr)
	if err != nil {
		panic(err)
	}

	store, err := storage.NewLearningStore(l, db)
	if err != nil {
		panic(err)
	}

	srv := http.NewServer(l, *addrPort, store)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	log.Printf("started server on port %s", *addrPort)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		cancel()
		log.Println("teardown complete")
	}()

	if err := srv.ShutDown(ctx); err != nil {
		panic(err)
	}
	log.Println("gracefully shutdown")
}

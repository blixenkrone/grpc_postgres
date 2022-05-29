package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/blixenkrone/lea/server"
	"github.com/blixenkrone/lea/storage"
	"github.com/blixenkrone/lea/storage/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	env  = flag.String("env", "local", "Environment")
	port = flag.String("port", ":9090", "The server port")
)

func main() {
	if err := godotenv.Load(fmt.Sprintf("%s.env", *env)); err != nil {
		panic(err)
	}

	l := logrus.New()

	pgConnStr, ok := os.LookupEnv("PG_CONN_STRING")
	if !ok {
		panic("no connstr found")
	}
	pgdb, err := postgres.NewFromConnectionString(pgConnStr)
	if err != nil {
		panic(err)
	}
	learningsDB, err := storage.NewLearningStore(l, pgdb)
	if err != nil {
		panic(err)
	}

	if err := learningsDB.MigrateUp("../../../storage/postgres/migrations"); err != nil {
		panic(err)
	}

	srv := server.NewGRPC(l, learningsDB)
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		l.Infof("starting gRPC on port: %s", *port)
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()
	<-done
	srv.GracefulStop()
	log.Println("gracefully shutdown")

}

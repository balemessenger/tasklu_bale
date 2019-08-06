package cmd

import (
	"fmt"
	"taskulu/api/grpc"

	"taskulu/api/http"
	"taskulu/internal"
	"taskulu/internal/server"
	"taskulu/pkg"

	"taskulu/internal/postgres"
)

func initialize() *pkg.Logger {
	fmt.Println("taskulu build version:", pkg.BuildVersion)
	fmt.Println("taskulu build time:", pkg.BuildTime)
	conf := internal.NewConfig("")
	log := pkg.NewLog(conf.Log.Level)

	db := postgres.New(log, postgres.Option{
		Host: conf.Postgres.Host,
		User: conf.Postgres.User,
		Pass: conf.Postgres.Pass,
		Db:   conf.Postgres.DB,
	})

	grpc.New(log, grpc.Option{
		Address: conf.Endpoints.Grpc.Address,
	})

	http.New(
		log,
		http.Option{
			Address: conf.Endpoints.Http.Address,
			User:    conf.Endpoints.Http.User,
			Pass:    conf.Endpoints.Http.Pass,
		})

	pkg.NewPrometheus(log, conf.Prometheus.Port)

	//Initialize main logic
	internal.NewExample(log, db).Start(conf.Core.WorkPoolSize)

	return log
}

func Main() {
	log := initialize()
	log.Info("Hello taskulu")
	server.New().Run()
	pkg.Signal.Wait()
}

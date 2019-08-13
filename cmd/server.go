package cmd

import (
	"fmt"
	"taskulu/api/http"
	"taskulu/internal"
	"taskulu/internal/server"
	"taskulu/pkg"
)

func initialize() *pkg.Logger {
	fmt.Println("taskulu build version:", pkg.BuildVersion)
	fmt.Println("taskulu build time:", pkg.BuildTime)
	conf := internal.NewConfig("")
	log := pkg.NewLog(conf.Log.Level)

	http.New(
		log,
		http.Option{
			Address: conf.Endpoints.Http.Address,
			User:    conf.Endpoints.Http.User,
			Pass:    conf.Endpoints.Http.Pass,
		})

	pkg.NewPrometheus(log, conf.Prometheus.Port)

	internal.RunIntegration(log, "672ba3ce56037687f59fc746bf32f60581d8c551d5ead7aa098697021443700e", "5d088afd56ad6678a4df44dc", "مولانا")
	internal.RunIntegration(log, "0ad8c705c806b74e7737d091f446d960366dccde655b879b95e2e861937d7e71", "5d088afd56ad6678a4df44dc", "گوشک")

	return log
}

func Main() {
	log := initialize()
	log.Info("Hello taskulu")
	server.New().Run()
	pkg.Signal.Wait()
}

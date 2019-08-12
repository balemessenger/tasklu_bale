package cmd

import (
	"fmt"
	"time"

	"taskulu/api/http"
	"taskulu/internal"
	"taskulu/internal/server"
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
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

	bale := internal.NewBale("https://api.bale.ai", "672ba3ce56037687f59fc746bf32f60581d8c551d5ead7aa098697021443700e")

	task := taskulu.New(log, taskulu.Option{
		BaseUrl:  "https://taskulu.com",
		Username: "amsjavan",
		Password: "0",
	})

	activity := internal.NewActivity(log, task, time.Now())
	integration := internal.NewBaleIntegration(log, bale, activity)

	integration.Run()
}

func Main() {
	log := initialize()
	log.Info("Hello taskulu")
	server.New().Run()
	pkg.Signal.Wait()
}

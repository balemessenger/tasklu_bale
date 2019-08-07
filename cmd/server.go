package cmd

import (
	"fmt"
	"taskulu/api/grpc"
	"time"

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

	bale := internal.NewBale("https://api.bale.ai", "672ba3ce56037687f59fc746bf32f60581d8c551d5ead7aa098697021443700e")

	taskulu := internal.NewTaskulu("https://taskulu.com")

	date := time.Now()

	for {
		err, body := taskulu.GetActivities("AhlIDxPy5cphIr_O9DPUQ-7jetC3y8wpsCzf9m9TBeJ6IYA9cwdWO0dVn7znYU8z8Lx0wXt_M41HD3FVMH1Wqqkwyvmk5gG_LygeonZepSn557299wY31pwlnr802HsS", "8cdf91ddd058682d80163b7ecb93116a", "5a8d1fff56ad660b0dd0d343")
		if err != nil {
			log.Error(err)
		}
		t := time.Unix(int64(body.Data[0].CreatedAt), 0)
		if t.After(date) {
			message := "تسک "
			message += body.Data[0].Content.Keys[0].Value
			message += " از وضعیت "
			message += body.Data[0].Content.Keys[1].Value
			message += " به وضعیت "
			message += body.Data[0].Content.Keys[2].Value
			message += " تغییر کرد."
			err = bale.Send(message)
			if err != nil {
				log.Error("BaleHook error::", err)
			}
			date = t
		}
		time.Sleep(time.Second)
	}
}

func Main() {
	log := initialize()
	log.Info("Hello taskulu")
	server.New().Run()
	pkg.Signal.Wait()
}

package cmd

import (
	"context"
	"fmt"
	context2 "golang.org/x/net/context"
	"taskulu/api/http"
	"taskulu/internal"
	"taskulu/internal/bot"
	"taskulu/internal/postgres"
	"taskulu/internal/server"
	"taskulu/pkg"
	"time"
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

	pg := postgres.New(log, postgres.Option{
		Host: "127.0.0.1",
		Port: "5432",
		User: "nasim",
		Pass: "nasim",
		Db:   "taskulu",
	})

	b := bot.New(log, pg, "512522922:3a615125d2d088f1768c9fb29acb9d7ff5127ad8")
	b.Run()

	//Run all users
	readUsers := make(map[int]context2.CancelFunc)

	date := time.Unix(0, 0)

	for {
		credentials, err := pg.GetAllUserAuth(date)
		if err != nil {
			log.Error(err)
		}
		for _, cred := range credentials {
			if cancel, ok := readUsers[cred.UserId]; ok {
				cancel()
			}
			ctx, cancel := context.WithCancel(context.Background())
			internal.RunNotification(ctx, log, b, cred.UserId, cred.Username, cred.Password)
			if cred.UpdatedAt.After(date) {
				date = cred.UpdatedAt
			}
			readUsers[cred.UserId] = cancel
		}
		time.Sleep(5 * time.Second)
	}

	////تندر بله - گوشک
	//internal.RunIntegration(log, "672ba3ce56037687f59fc746bf32f60581d8c551d5ead7aa098697021443700e", "5d088afd56ad6678a4df44dc", "گوشک")
	//
	////تندر بله - مولانا
	//internal.RunIntegration(log, "663740aa141d26f57eab0d1ef75078652378edbf5479f981af7f98c0ba6abbf3", "5d088afd56ad6678a4df44dc", "مولانا")
	//
	////Bale Ticket - مولانا و فروغ
	//internal.RunIntegration(log, "6bfe43fe49ae3ee54ab44fbe716036bd6ca174ccdc0bce1875108a6b3d81fd69", "5a8d1fff56ad660b0dd0d343", "فروغ")
	////Bale Ticket - مولانا و سعدی
	//internal.RunIntegration(log, "da6b3be7ecf67e874816b4da5e21feec62fd9ad982df480074ab5d60c1a2d9ab", "5a8d1fff56ad660b0dd0d343", "سعدی")
	////Bale Ticket - مولانا و حافظ
	//internal.RunIntegration(log, "5bd7284e0698891932e2fbbfdc2b72982d5fd023143aa5bce4269d9e9f1554da", "5a8d1fff56ad660b0dd0d343", "حافظ")
	////Bale Ticket - مولانا و پروین
	//internal.RunIntegration(log, "ad5c8d8361e773ecb795b9e4d6a72370f0efaa4dd3e6f34ae15c4c2a3fc26f29", "5a8d1fff56ad660b0dd0d343", "پروین")
	////Bale Ticket - مولانا و قیصر
	//internal.RunIntegration(log, "0d2d9812abc55235e6bd3e292b6f367042c5621078bf8e6cbf982a25192ef251", "5a8d1fff56ad660b0dd0d343", "قیصر")
	////Bale Ticket - مولانا و فردوسی
	//internal.RunIntegration(log, "5e196aae3dd8f47f6a978b69e30764585fb06a71e0e60dddb541081723367c39", "5a8d1fff56ad660b0dd0d343", "فردوسی")
	////Bale Ticket - مولانا و شهریار
	//internal.RunIntegration(log, "2cc1ce990d75b2d58118b2db8eb37f2d7066f1cb25e71148b5675ebb71056963", "5a8d1fff56ad660b0dd0d343", "شهریار")
	////Bale Ticket - مولانا و بازاریابی
	//internal.RunIntegration(log, "470825625ae768c33ed954622152b596181e5b0fe376314712b013c9308b759b", "5a8d1fff56ad660b0dd0d343", "بازاریابی")

	return log
}

func Main() {
	log := initialize()
	log.Info("Hello taskulu")
	server.New().Run()
	pkg.Signal.Wait()
}

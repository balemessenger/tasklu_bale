package test

import (
	"taskulu/internal"
	"taskulu/testkit/mock"

	"math/rand"
	"os"
	"taskulu/api/http"
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"taskulu/testkit"
	"testing"
	"time"
)

var (
	Conf        *internal.Config
	task        *taskulu.Client
	integration *internal.BaleIntegration
)

func setup() {
	rand.Seed(time.Now().Unix())
	Conf = testkit.InitTestConfig("config.yaml")
	log := pkg.NewLog("DEBUG")

	http.New(
		log,
		http.Option{
			Address: Conf.Endpoints.Http.Address,
			User:    Conf.Endpoints.Http.User,
			Pass:    Conf.Endpoints.Http.Pass,
		})

	mock.New(log, mock.Option{
		Address: "127.0.0.1:12346",
		User:    "test",
		Pass:    "test",
	})

	task = taskulu.New(log, taskulu.Option{
		"http://127.0.0.1:12346",
		"test",
		"test",
	})

	bale := internal.NewBale("http://127.0.0.1:12346", "")

	activity := internal.NewActivity(log, task, time.Unix(1565091220, 0))
	integration = internal.NewBaleIntegration(log, bale, activity)

	time.Sleep(4000 * time.Millisecond)
}

func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	r := m.Run()
	teardown()
	os.Exit(r)
}

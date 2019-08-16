package test

import (
	"taskulu/internal"
	postgres2 "taskulu/internal/postgres"
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
	log         *pkg.Logger
	postgres    *postgres2.Database
	Conf        *internal.Config
	task        *taskulu.Client
	integration *internal.BaleIntegration
)

func setup() {
	rand.Seed(time.Now().Unix())
	Conf = testkit.InitTestConfig("config.yaml")
	log = pkg.NewLog("DEBUG")

	http.New(
		log,
		http.Option{
			Address: Conf.Endpoints.Http.Address,
			User:    Conf.Endpoints.Http.User,
			Pass:    Conf.Endpoints.Http.Pass,
		})

	postgres = postgres2.New(log, postgres2.Option{
		Host: "127.0.0.1",
		Port: "5432",
		User: "taskulu",
		Pass: "taskulu",
		Db:   "taskulu",
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

	projectId := "12346"
	sheetName := "SampleSheet"
	sheet := internal.NewSheet(log, task)
	activity := internal.NewActivity(log, task, sheet, time.Unix(1565091220, 0))
	integration = internal.NewBaleIntegration(log, bale, activity, projectId, sheetName)

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

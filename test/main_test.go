package test

import (
	"taskulu/internal"
	"taskulu/testkit/mock"

	"math/rand"
	"os"
	"taskulu/api/http"
	"taskulu/pkg"
	"taskulu/testkit"
	"testing"
	"time"
)

var Conf *internal.Config

var taskulu *internal.TaskuluClient

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

	taskulu = internal.NewTaskulu(log, internal.Option{
		"http://127.0.0.1:12346",
		"test",
		"test",
	})

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

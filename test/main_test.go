package test

import (
	grpc2 "taskulu/api/grpc"
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

var mockBaleHook *mock.FakeServer
var mockTaskulu *mock.FakeServer

func setup() {
	rand.Seed(time.Now().Unix())
	Conf = testkit.InitTestConfig("config.yaml")
	log := pkg.NewLog("DEBUG")

	grpc2.New(log, grpc2.Option{
		Address: Conf.Endpoints.Grpc.Address,
	})

	testkit.GetGrpcClient().Initialize(Conf.Endpoints.Grpc.Address)

	http.New(
		log,
		http.Option{
			Address: Conf.Endpoints.Http.Address,
			User:    Conf.Endpoints.Http.User,
			Pass:    Conf.Endpoints.Http.Pass,
		})

	mockBaleHook = mock.NewMockServer("127.0.0.1:12345", "/v1/webhooks/")
	mockBaleHook.Start()

	mockTaskulu = mock.NewMockServer("127.0.0.1:12346", "/api/v1/projects/123456/activities")
	mockTaskulu.Start()

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

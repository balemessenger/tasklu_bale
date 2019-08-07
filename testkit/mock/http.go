package mock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FakeServer struct {
	address  string
	path     string
	handlers []func(http.ResponseWriter, []byte)
}

func NewMockServer(address string, path string) *FakeServer {
	return &FakeServer{address, path, nil}
}

func (f *FakeServer) Start() {
	go f.run()
}

func (f *FakeServer) AddHandler(handler func(http.ResponseWriter, []byte)) {
	f.handlers = append(f.handlers, handler)
}

func (f *FakeServer) run() {
	http.HandleFunc(f.path, func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 404)
			return
		}
		if f.handlers == nil {
			return
		}
		for _, h := range f.handlers {
			h(w, body)
		}
	})

	log.Fatal(http.ListenAndServe(f.address, nil))
}

package internal

import (
	"fmt"
	"net/http"
	"taskulu/internal/processor"
	"taskulu/internal/repositories"
	"taskulu/pkg"
)

type ExampleProcessor struct {
	log *pkg.Logger
	db  repositories.Database
	processor.Processor
	exampleChannel chan string
}

func NewExample(log *pkg.Logger, db repositories.Database) *ExampleProcessor {
	return &ExampleProcessor{
		log:            log,
		db:             db,
		Processor:      processor.New(),
		exampleChannel: make(chan string)}
}

func (g *ExampleProcessor) Start(size int) {
	g.RunPool(g.Processor, size)
}

func (g *ExampleProcessor) Tell(envelop string) {
	g.exampleChannel <- envelop
}

func (g *ExampleProcessor) Worker() {
	resp, err := http.Get("")
	if err != nil {
		g.log.Error(err)
	}
	if resp.StatusCode == 200 {
		fmt.Print(resp.Body)
	}
}

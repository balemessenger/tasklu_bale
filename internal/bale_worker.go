package internal

import (
	"taskulu/internal/processor"
	"taskulu/internal/repositories"
	"taskulu/pkg"
)

type BaleWorkerProcessor struct {
	log *pkg.Logger
	db  repositories.Database
	processor.Processor
	exampleChannel chan string
}

func NewBaleWorker(log *pkg.Logger, db repositories.Database) *ExampleProcessor {
	return &ExampleProcessor{
		log:            log,
		db:             db,
		Processor:      processor.New(),
		exampleChannel: make(chan string)}
}

func (g *BaleWorkerProcessor) Start(size int) {
	g.RunPool(g.Processor, size)
}

func (g *BaleWorkerProcessor) Tell(envelop string) {
	g.exampleChannel <- envelop
}

func (g *BaleWorkerProcessor) Worker() {
	for {
		envelop := <-g.exampleChannel
		g.process(envelop)
	}
}

func (g *BaleWorkerProcessor) process(envelop string) {
	//Write your login here

	g.log.Info("process")
}

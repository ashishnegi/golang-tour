package main

import (
	"fmt"
	"time"
)

type event struct {
	pid int
	tid int
	msg string
}

type csvevent struct {
	pid    int
	tid    int
	msg    string
	lineno int
}

func readEtlFiles(file string) (stream chan string) {
	go func() {
		for {
			fmt.Println("sending event.")
			stream <- "10 100 " + file + ":etl_message"
		}
	}()
	return stream
}

func parseEtlFiles(msgChan chan string) (events chan event) {
	go func() {
		for {
			fmt.Println("looking for msg")
			msg := <-msgChan
			fmt.Println("sending parsed etl.")
			events <- event{pid: 10, tid: 100, msg: msg}
		}
	}()
	return events
}

func publishEtlEvents(eventChan chan event, listeners ...chan event) {
	for {
		event := <-eventChan
		for i := range listeners {
			listeners[i] <- event
		}
	}
}

func etlToCsv(e event) csvevent {
	return csvevent{pid: e.pid, tid: e.tid, msg: e.msg, lineno: 10}
}

func publishEtlToCsvEvents(eventChan chan event, csvListeners ...chan csvevent) {
	for {
		event := <-eventChan
		csvEvent := etlToCsv(event)
		for i := range csvListeners {
			csvListeners[i] <- csvEvent
		}
	}
}

func main() {
	fabricFileCh := readEtlFiles("fabric_events.etl")
	ktlFileCh := readEtlFiles("ktl_events.etl")
	leaseFileCh := readEtlFiles("lease_events.etl")

	fabricEventsCh := parseEtlFiles(fabricFileCh)
	ktlEventsCh := parseEtlFiles(ktlFileCh)
	leaseEventsCh := parseEtlFiles(leaseFileCh)

	azureUploaderCh := make(chan event)
	inMemoryProducerCh := make(chan event)
	go publishEtlEvents(fabricEventsCh, azureUploaderCh, inMemoryProducerCh)
	go publishEtlEvents(ktlEventsCh, azureUploaderCh, inMemoryProducerCh)

	azureCsvUploaderCh := make(chan csvevent)
	go publishEtlToCsvEvents(leaseEventsCh, azureCsvUploaderCh)

	printEvtToScreen := func(name string, achan chan event) {
		e := <-achan
		fmt.Printf("%s : %+v\n", name, e)
	}

	printCsvToScreen := func(name string, achan chan csvevent) {
		e := <-achan
		fmt.Printf("%s : %+v\n", name, e)
	}

	go printEvtToScreen("azureuploader", azureUploaderCh)
	go printEvtToScreen("inmemoryproducer", inMemoryProducerCh)
	go printCsvToScreen("azurecsvuploader", azureCsvUploaderCh)

	stopCh := make(chan bool)
	go func() {
		time.Sleep(3000 * time.Millisecond)
		stopCh <- true
	}()

	<-stopCh
}

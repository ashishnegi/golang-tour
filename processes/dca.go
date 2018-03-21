package main

import (
	"fmt"
	"time"
)

type event struct {
	pid         int
	tid         int
	msg         string
	processedBy int
}

type csvevent struct {
	pid    int
	tid    int
	msg    string
	lineno int
}

func readEtlFiles(file string) chan string {
	stream := make(chan string)
	go func() {
		for {
			fmt.Printf("read one line from file : %s\n", file)
			time.Sleep(100 * time.Millisecond)
			stream <- "10 100 " + file + ":etl_message"
		}
	}()
	return stream
}

func parseEtlFiles(msgChan chan string) chan event {
	events := make(chan event)
	go func() {
		for {
			msg := <-msgChan
			fmt.Println("Parsed msg to etl")
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

func fanOut(in chan event, outs ...chan event) {
	for i := range outs {
		go func(i int) {
			for {
				msg, ok := <-in
				msg.processedBy = i
				if ok {
					outs[i] <- msg
				} else {
					break
				}
			}
		}(i)
	}
}

func fanIn(out chan event, ins ...chan event) {
	for i := range ins {
		go func(in chan event) {
			for {
				msg, ok := <-in
				if ok {
					out <- msg
				} else {
					break
				}
			}
		}(ins[i])
	}
}

// so, we find that parsing seems to be taking lot of time.
// and fabric_events has most of the events.
// we need to double up processing for fabric_events.
func main() {
	fabricFileCh := readEtlFiles("fabric_events.etl")
	ktlFileCh := readEtlFiles("ktl_events.etl")
	leaseFileCh := readEtlFiles("lease_events.etl")

	fabricEventsOutCh1 := make(chan event)
	fabricEventsOutCh2 := make(chan event)
	fabricEventsInCh := parseEtlFiles(fabricFileCh)
	fanOut(fabricEventsInCh, fabricEventsOutCh1, fabricEventsOutCh2)
	fabricEventsCh := make(chan event)
	fanIn(fabricEventsCh, fabricEventsOutCh1, fabricEventsOutCh2)

	ktlEventsCh := parseEtlFiles(ktlFileCh)
	leaseEventsCh := parseEtlFiles(leaseFileCh)

	azureUploaderCh := make(chan event)
	inMemoryProducerCh := make(chan event)
	go publishEtlEvents(fabricEventsCh, azureUploaderCh, inMemoryProducerCh)
	go publishEtlEvents(ktlEventsCh, azureUploaderCh, inMemoryProducerCh)

	azureCsvUploaderCh := make(chan csvevent)
	go publishEtlToCsvEvents(leaseEventsCh, azureCsvUploaderCh)

	printEvtToScreen := func(name string, achan chan event) {
		for {
			e := <-achan
			fmt.Printf("%s : %+v\n", name, e)
		}
	}

	printCsvToScreen := func(name string, achan chan csvevent) {
		for {
			e := <-achan
			fmt.Printf("%s : %+v\n", name, e)
		}
	}

	go printEvtToScreen("azureuploader", azureUploaderCh)
	go printEvtToScreen("inmemoryproducer", inMemoryProducerCh)
	go printCsvToScreen("azurecsvuploader", azureCsvUploaderCh)

	time.Sleep(1000 * time.Millisecond)
}

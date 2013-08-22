package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type GraphiteReporter struct {
	Config GraphiteConfig
}

func NewGraphiteReporter(conf GraphiteConfig) (h *GraphiteReporter) {
	return &GraphiteReporter{Config: conf}
}

func (self *GraphiteReporter) ReportHealth(h *Health) {
	hmap := h.Map()
	data := ""
	now := time.Now()
	ts := now.Unix()

	for k, v := range hmap {
		data += fmt.Sprintf("%s%s%s %v %v\n", self.Config.Prefix, k, self.Config.Postfix, v, ts)
	}

	addr, err := net.ResolveTCPAddr("tcp", self.Config.LineRec)
	if err != nil {
		log.Println("Graphite: Cannot resolve address: ", err.Error())
		return
	}
	// open up a connection each time. this dismisses the complexity of keeping
	// connection state maintained at all times.
	// IMHO no specialized need in this case for keeping a TCP conn. open.
	conn, err := net.DialTCP("tcp", nil, addr)
	defer conn.Close()

	if err != nil {
		log.Println("Graphite: Cannot connect: ", err.Error())
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		log.Println("Graphite: Cannot write data on connection: ", err.Error())
	}
}

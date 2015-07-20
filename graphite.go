package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"net"
)

type Graphite struct {
	net.Conn
	prefix string
}

func NewGraphite(address string, prefix string) (*Graphite, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, err
	}

	return &Graphite{conn, prefix}, nil
}

// send sends an object count to Graphite.
func (g *Graphite) Send(val int) {
	line := fmt.Sprintf("%s %d", g.prefix, val)
	logrus.Debugf("Sending: %v", line)
	_, _ = fmt.Fprintf(g.Conn, "%s\n", line)
}

// Close the UDP connection.
// TODO: Is this even necesssary?
func (g *Graphite) Close() error {
	if g.Conn != nil {
		return g.Conn.Close()
	}

	return nil
}

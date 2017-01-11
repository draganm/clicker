package server

import (
	"net"

	"github.com/draganm/zathras/topic"
)

var Topic *topic.Topic

func Serve(addr string, t *topic.Topic) error {
	Topic = t
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	c, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer c.Close()

	buf := make([]byte, 65536)

	for {
		r, err := c.Read(buf)
		if err != nil {
			return err
		}

		_, err = t.WriteEvent(buf[:r])
		if err != nil {
			return err
		}
	}

}

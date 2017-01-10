package server

import (
	"net"

	"github.com/draganm/zathras/topic"
)

func Serve(addr string, t *topic.Topic) error {
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

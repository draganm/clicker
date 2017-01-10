package main

import (
	"os"

	"github.com/draganm/clicker/proxy"
	"github.com/draganm/clicker/server"
	"github.com/draganm/zathras/topic"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			&cli.Command{
				Name: "server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "server-bind",
						Usage: "host:port where servier will listen for events. If it should listen to any host use ':port'",
						Value: ":5555",
					},
					&cli.StringFlag{
						Name:  "web-bind",
						Usage: "host:port where servier will listen for http clients. If it should listen to any host use ':port'",
						Value: ":9999",
					},
					&cli.StringFlag{
						Name:  "data-dir",
						Usage: "Directory where event data will be stored",
						Value: "events/",
					},
				},
				Action: func(c *cli.Context) error {
					dir := c.String("data-dir")
					topic, err := topic.New(dir, 20*1024*1024)
					if err != nil {
						return err
					}
					return server.Serve(c.String("server-bind"), topic)
				},
			},
			&cli.Command{
				Name: "proxy",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "bind, b",
						Value: ":8081",
						Usage: "host:port where proxy will bind. If it should listen to any host use ':port'",
					},
					&cli.StringFlag{
						Name:  "backend-url",
						Usage: "url of the backend for which the traffic should be proxied",
						Value: "http://localhost:8080",
					},
					&cli.StringFlag{
						Name:  "server-addr",
						Usage: "Address of the server that collects the events",
						Value: "localhost:5555",
					},
				},
				Action: func(c *cli.Context) error {
					return proxy.Proxy(c.String("bind"), c.String("backend-url"), c.String("server-addr"))
				},
			},
		},
	}
	app.Run(os.Args)
}

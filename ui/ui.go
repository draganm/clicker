package ui

import reactor "github.com/draganm/go-reactor"

func Serve(addr string) error {
	reactor := reactor.New()
	reactor.AddScreen("/", IndexFactory)
	reactor.Serve(addr)
	return nil
}

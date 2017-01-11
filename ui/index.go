package ui

import (
	"fmt"
	"sync"

	"github.com/draganm/clicker/query"
	"github.com/draganm/clicker/server"
	reactor "github.com/draganm/go-reactor"
)

var indexUI = reactor.MustParseDisplayModel(`
  <dl>
    <dt>Reuquests #</dt>
    <dd id="requestCount">...</dd>
    <dt>Bytes Received</dt>
    <dd id="bytesReceived">...</dd>
    <dt>Bytes Sent</dt>
    <dd id="bytesSent">...</dd>
    <dt>Errors</dt>
    <dd id="errors">...</dd>
  </dl>
`)

func IndexFactory(ctx reactor.ScreenContext) reactor.Screen {
	return &Index{
		ctx: ctx,
	}
}

type Index struct {
	sync.Mutex
	ctx     reactor.ScreenContext
	soFar   *query.SoFar
	current query.WebStatsBucket
}

func (i *Index) Mount() {
	i.soFar = query.NewSoFar(server.Topic)
	i.soFar.AddUpdateListener(i.onWebStatsBucket)
}

func (i *Index) onWebStatsBucket(b query.WebStatsBucket) {
	i.Lock()
	defer i.Unlock()
	i.current = b
	i.render()
}

func (i *Index) OnUserEvent(evt *reactor.UserEvent) {
}

func (i *Index) render() {
	ui := indexUI.DeepCopy()
	ui.SetElementText("requestCount", fmt.Sprintf("%d", i.current.Requests))
	ui.SetElementText("bytesReceived", fmt.Sprintf("%d", i.current.BytesReceived))
	ui.SetElementText("bytesSent", fmt.Sprintf("%d", i.current.BytesSent))
	ui.SetElementText("errors", fmt.Sprintf("%d", i.current.Errors))
	// bytesReceived
	i.ctx.UpdateScreen(&reactor.DisplayUpdate{
		Model: WithNavigation(ui),
	})
}

func (i *Index) Unmount() {
	i.soFar.RemoveUpdateListener(i.onWebStatsBucket)
	i.soFar.Close()
}

var navigationUI = reactor.MustParseDisplayModel(`
  <div>
  	<bs.Navbar bool:fluid="true">
  		<bs.Navbar.Header>
  			<bs.Navbar.Brand>
  				<a href="#" className="navbar-brand">Clicker</a>
  			</bs.Navbar.Brand>
  		</bs.Navbar.Header>
  	</bs.Navbar>
  	<bs.Grid bool:fluid="true">
			<bs.Row>
				<bs.Col int:mdOffset="1" int:md="10" int:smOffset="0" int:sm="12">
					<div id="content" className="container"></div>
				</bs.Col>
			</bs.Row>
  	</bs.Grid>
  </div>
`)

func WithNavigation(content *reactor.DisplayModel) *reactor.DisplayModel {
	view := navigationUI.DeepCopy()
	view.ReplaceChild("content", content)
	return view
}

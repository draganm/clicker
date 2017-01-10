package integration_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/draganm/clicker/comm"
	"github.com/draganm/clicker/proxy"
	"github.com/draganm/clicker/server"
	"github.com/draganm/zathras/topic"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func startHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	go http.ListenAndServe("localhost:8081", mux)
	for {
		_, err := http.Get("http://localhost:8081")
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(time.Millisecond * 10)
	}

}

func startClickerProxy() {

	go func() {
		err := proxy.Proxy(":8080", "http://localhost:8081", "localhost:3333")
		if err != nil {
			panic(err)
		}
	}()

	for {
		_, err := http.Get("http://localhost:8080")
		if err == nil {
			break
		}
		time.Sleep(time.Millisecond * 10)
	}

}

var _ = AfterSuite(func() {
	if t != nil {
		Expect(t.Close()).To(Succeed())
	}
	Expect(os.RemoveAll(topicDir)).To(Succeed())
})

var t *topic.Topic
var topicDir string

func startClickerServer() {
	var err error
	topicDir, err = ioutil.TempDir("", "")
	Expect(err).ToNot(HaveOccurred())

	t, err = topic.New(topicDir, 1024*1024*10)
	Expect(err).ToNot(HaveOccurred())

	// proxy.Proxy()
	go server.Serve(":3333", t)

}

var _ = BeforeSuite(func(done Done) {
	startHTTPServer()
	startClickerProxy()
	startClickerServer()
	close(done)
}, 3.0)

var _ = Describe("Logging", func() {

	Context("When a POST request passes the clicker proxy", func() {
		var response *http.Response
		BeforeEach(func() {
			var err error
			request, err := http.NewRequest("POST", "http://localhost:8080/test1", strings.NewReader("testRequest"))
			Expect(err).ShouldNot(HaveOccurred())
			response, err = http.DefaultClient.Do(request)
			Expect(err).ShouldNot(HaveOccurred())
		})

		var s <-chan topic.Event
		var c chan interface{}

		BeforeEach(func() {
			s, c = t.Subscribe(t.LastID())
		})

		AfterEach(func() {
			close(c)
		})

		It("should receive 200 response code", func() {
			Expect(response.StatusCode).To(Equal(200))
		})

		It("Should log event", func(done Done) {

			{
				d := <-s
				evt, err := comm.Decode(d.Data)
				Expect(err).ToNot(HaveOccurred())
				Expect(evt.Method).To(Equal("POST"))
				Expect(evt.RequestURI).To(Equal("/test1"))
				Expect(evt.Time).ToNot(Equal(time.Time{}))
				Expect(evt.UUID).ToNot(BeEmpty())
				Expect(evt.RequestHeader).To(HaveKey("User-Agent"))
				Expect(evt.ResponseHeader).To(HaveKey("Content-Type"))
				Expect(string(evt.CapturedResponseBody)).To(Equal("test"))
				Expect(string(evt.CapturedRequestBody)).To(Equal("testRequest"))
				Expect(evt.StatusCode).To(Equal(200))
				Expect(evt.BytesRead).To(Equal(11))
				Expect(evt.BytesWritten).To(Equal(4))
				Expect(evt.LastByteReadAt).ToNot(Equal(time.Time{}))
				Expect(evt.LastByteWrittenAt).ToNot(Equal(time.Time{}))

			}

			close(done)
		})

	})

	Context("When a GET request passes the clicker proxy", func() {
		var response *http.Response
		BeforeEach(func() {
			var err error
			response, err = http.Get("http://localhost:8080/test1")
			Expect(err).ShouldNot(HaveOccurred())
		})

		var s <-chan topic.Event
		var c chan interface{}

		BeforeEach(func() {
			s, c = t.Subscribe(t.LastID())
		})

		AfterEach(func() {
			close(c)
		})

		It("should receive 200 response code", func() {
			Expect(response.StatusCode).To(Equal(200))
		})

		It("Should log event", func(done Done) {

			{
				d := <-s
				evt, err := comm.Decode(d.Data)
				Expect(err).ToNot(HaveOccurred())
				Expect(evt.Method).To(Equal("GET"))
				Expect(evt.RequestURI).To(Equal("/test1"))
				Expect(evt.Time).ToNot(Equal(time.Time{}))
				Expect(evt.UUID).ToNot(BeEmpty())
				Expect(evt.RequestHeader).To(HaveKey("User-Agent"))
				Expect(evt.ResponseHeader).To(HaveKey("Content-Type"))
				Expect(string(evt.CapturedResponseBody)).To(Equal("test"))
				Expect(string(evt.CapturedRequestBody)).To(Equal(""))
				Expect(evt.StatusCode).To(Equal(200))
				Expect(evt.BytesRead).To(Equal(0))
				Expect(evt.BytesWritten).To(Equal(4))
				Expect(evt.LastByteReadAt).ToNot(Equal(time.Time{}))
				Expect(evt.LastByteWrittenAt).ToNot(Equal(time.Time{}))

			}

			close(done)
		})

	})

})

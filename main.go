package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var (
	BestServer       *Server
	LeastConnections uint = 0
	FirstURL, _           = url.Parse("http://127.0.0.1:8081")
	SecondURL, _          = url.Parse("http://127.0.0.1:8082")
)

type Server struct {
	URL         url.URL
	Connections uint
	Alive       bool
	Mutex       sync.Mutex
}
type Servers []*Server

func (servers Servers) ChooseServer() (*Server, error) {
	for _, server := range servers {
		server.Mutex.Lock()
		if resp, err := http.Get(server.URL.String()); err != nil || resp.StatusCode >= 500 {
			server.Alive = false
			return BestServer, err
		} else {
			server.Alive = true
		}
		if (server.Connections < LeastConnections || LeastConnections == 0) && server.Alive {
			LeastConnections = server.Connections
			BestServer = server
		}
		if BestServer != nil {
			server.Connections++
		}
		server.Mutex.Unlock()
	}
	return BestServer, nil
}
func HandleRequest(w http.ResponseWriter, r *http.Request, servers Servers) {
	server, _ := servers.ChooseServer()
	if server == nil {
		fmt.Println("Currently there is no available servers there.")
		return
	}
	defer func() {
		server.Mutex.Lock()
		server.Connections--
		server.Mutex.Unlock()
	}()
	proxy := httputil.NewSingleHostReverseProxy(&server.URL)
	proxy.ServeHTTP(w, r)
}
func main() {
	servers := Servers{{URL: *FirstURL}, {URL: *SecondURL}}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		HandleRequest(writer, request, servers)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

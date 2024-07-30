<h1>Simple Load Balancer in Go</h1>

**Русская версия:** [здесь](https://github.com/An-another-Artem/simple-load-balancer-in-go-ru)

**What is a Load Balancer?**

A load balancer is a program that distributes client requests across multiple servers to prevent overloading any single server. This helps ensure high availability and performance for your application.

**Load Balancing Algorithms**

There are several load balancing algorithms, each with its own advantages and disadvantages. Here are three common ones:

* **Round Robin (including Weighted Round Robin and Sticky Round Robin)**: This algorithm distributes requests evenly among all available servers. Weighted Round Robin assigns more weight to servers with higher capacity, while Sticky Round Robin directs requests from the same client to the same server for session consistency.
* **Least Connections (Currently Implemented)**: This algorithm tracks the number of active connections on each server and assigns new requests to the server with the fewest connections, aiming for even load distribution.
* **Least Time**: This algorithm measures the response time of each server and directs requests to the server with the fastest response time. 

**Round Robin Algorithm**

<img src="https://www.jscape.com/hubfs/images/round_robin_algorithm-1.png">
In Round Robin, requests are cycled through all available servers, regardless of their current load.

**Least Connections Algorithm (Implemented)**

<img src="https://www.codereliant.io/content/images/2023/06/d1-1-1.png">
This algorithm tracks the number of **active connections** on each server and assigns new requests to the server with the **fewest connections**, aiming for **even load distribution**.

**Least Time Algorithm**

The Least Time algorithm measures the response time of each server and directs requests to the server with the fastest response time. It requires a method to periodically check server response times, which can add complexity.

**Note:**

I couldn't find a suitable image for the Least Time algorithm. You can find one yourself.

<h2>What is going next?</h2>

**What should we know:**

* The load balancer uses `httputil`'s package proxy to send client's requests to servers.
* It currently doesn't use config files or other mechanisms to add or change servers. You can add this feature as an improvement.

**Now we are ready to make the load balancer!**

The load balancer implements the "Least connections" algorithm. As mentioned earlier, it directs client requests to the server that has the fewest connections.

```Go
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
```
<h2>Thank you for reading this document. Have a great day!<h2>

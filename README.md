<h1>Simple Load Balancer in Go</h1>

**What is a Load Balancer?**

A load balancer is a program that distributes client requests across multiple servers to prevent overloading any single server. This helps ensure high availability and performance for your application.

**Load Balancing Algorithms**

There are several load balancing algorithms, each with its own advantages and disadvantages. Here are three common ones:

* **Round Robin (including Weighted Round Robin and Sticky Round Robin)**: This algorithm distributes requests evenly among all available servers. Weighted Round Robin assigns more weight to servers with higher capacity, while Sticky Round Robin directs requests from the same client to the same server for session consistency.
* **Least Connections**: This algorithm tracks the number of active connections on each server and assigns new requests to the server with the fewest connections, aiming for even load distribution. (**This is the algorithm we will be implementing in our Go program.**)
* **Least Time**: This algorithm measures the response time of each server and directs requests to the server with the fastest response time. 

**Round Robin Algorithm**

<img src="https://www.jscape.com/hubfs/images/round_robin_algorithm-1.png">
In Round Robin, requests are cycled through all available servers, regardless of their current load.

**Least Connections Algorithm**

<img src="https://www.codereliant.io/content/images/2023/06/d1-1-1.png">
The Least Connections algorithm keeps track of the number of active connections on each server and directs new requests to the server with the fewest connections, aiming for a more balanced load distribution.

**Least Time Algorithm**

The Least Time algorithm measures the response time of each server and directs requests to the server with the fastest response time. It requires a method to periodically check server response times, which can add complexity.

**Note:**

I couldn't find a suitable image for the Least Time algorithm. You can find one yourself and update the link accordingly.

<h2>What is going next?</h2>

**What should we know**

* The load balancer will use httputil's package proxy to send client's request on server
* It won't use config files or something else to add or change servers, so you may add this feature
* You should give star on this repository as well

**Now we are ready to make the load balancer**

The load balancer will use "Least connections" algorithm and as i said upper our program will direct client's request to server that have the fewest connections.


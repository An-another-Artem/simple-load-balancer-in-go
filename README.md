<h1>Simple load balancer in go</h1>
First of all we should understand what is load balancer in total. <b>Load balancer<b> is a program that balances client requests so that the servers do not burn out.
Second of all to understand the load balancers we need to know their algoritms. In general, there are 3 large groups of them, they are:
<b><ul>Robin Round (it has several types such as Weighted Robin Round and Sticky Robin Round)</ul></b>
<b><ul>Least connections (we are going to use it in load balancer)</ul></b>
<b><ul>Least time</ul></b>
In <b>Robin Round</b> algoritm load balancer allocate requests to all available servers.
<img src="https://www.jscape.com/hubfs/images/round_robin_algorithm-1.png">
<b>Least connections</b> algoritm counts and stores client's connection to each server.
<img src="https://www.researchgate.net/publication/347808307/figure/fig2/AS:1000876902723590@1615639057440/Least-connection-algorithm.png">
<b>Least time</b> algoritm sends next request to server which has least ping to the client.
<i>well, i couldn't found the good image, so imagine in yourself :)</i>

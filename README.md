# Simple ping-client and pong-server in Go

## ping-client

Run as `ping-client [-port=8888] [-pong-service=pong-server.service.consul] [-dns=172.17.42.1:53]`.

`ping-client` is an http server which listens to `/`, performs GET request to a URL specified with `-pong-server` and returns received response.

## pong-server

Run as `pong-client [-port=8000]`.

`pong-server` is an http server which listens to `/` and returns hostname and network addresses of a serving host.


## Service discovery demo

- DNS config for Docker daemon inside of boot2docker VM is not necessarry, because we specify DNS server explicitly when we run ping-client.
- Be careful when starting manually a Consul node with exposed DNS port. For convinience it's better to bind it to 53 (default DNS port), e.g. `-p 53:53`.

### Basic expirement

3 nodes are defined in Vagrantfile and 3 corresponding scripts in `env` directory have all needed commands to start Consul servers, Registrator instances, pong-server and ping-client apps.

Both pong-server and ping-client Docker images can be built using `docker build -t <name> .` command executed in correspoding subdirectories.

## Resources

### Deployment

- [Run Go app with Docker](https://blog.golang.org/docker)
- [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07)
- [registrator](https://github.com/progrium/registrator)
- [docker-consul](https://github.com/progrium/docker-consul)
- [DNS Client Load Balancer](https://github.com/benschw/dns-clb-go)

### Learn more

- [DNS library in Go](https://github.com/miekg/dns)
- [WHY YOU DON'T NEED TO RUN SSHD IN YOUR DOCKER CONTAINERS](https://blog.docker.com/2014/06/why-you-dont-need-to-run-sshd-in-docker/)
- [Consul in production](https://www.digitalocean.com/community/tutorials/how-to-configure-consul-in-a-production-environment-on-ubuntu-14-04)
- [Using Consul with DNSMasq and other tools](http://johnhamelink.com/distributed-web-systems-with-consul-haproxy-and-envoy.html)

# Simple ping-client and pong-server in Go

## ping-client

Run as `ping-client [-port=8888] [-pong-server=http://pong:8000]`.

`ping-client` is an http server which listens to `/`, performs GET request to a URL specified with `-pong-server` and returns received response.

## pong-server

Run as `pong-client [-port=8000]`.

`pong-server` is an http server which listens to `/` and returns hostname and network addresses of a serving host.
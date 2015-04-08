#!/bin/sh

# main consul node

export HOSTNAME=node-01
PUBLIC_IP=192.168.33.10

docker run --name consul -h $HOSTNAME \
	-p $PUBLIC_IP:8300:8300 \
	-p $PUBLIC_IP:8301:8301 \
	-p $PUBLIC_IP:8301:8301/udp \
	-p $PUBLIC_IP:8302:8302 \
	-p $PUBLIC_IP:8302:8302/udp \
	-p $PUBLIC_IP:8400:8400 \
	-p $PUBLIC_IP:8500:8500 \
	-p 172.17.42.1:53:53/udp \
	-p $PUBLIC_IP:53:53/udp \
	-d -v /mnt:/data progrium/consul -server \
	-advertise $PUBLIC_IP \
	-bootstrap-expect 3 \
	-ui-dir /ui

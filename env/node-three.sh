#!/bin/sh

export HOSTNAME=node-three
PUBLIC_IP=192.168.33.12
JOIN_IP=192.168.33.10

# run consul server
docker run --name consul -h $HOSTNAME \
	-p $PUBLIC_IP:8300:8300 \
	-p $PUBLIC_IP:8301:8301 \
	-p $PUBLIC_IP:8301:8301/udp \
	-p $PUBLIC_IP:8302:8302 \
	-p $PUBLIC_IP:8302:8302/udp \
	-p $PUBLIC_IP:8400:8400 \
	-p $PUBLIC_IP:8500:8500 \
	-p 172.17.42.1:53:53/udp \
	-d -v /mnt:/data progrium/consul -server \
	-advertise $PUBLIC_IP \
	-join $JOIN_IP

# run registrator
docker run -d \
    -v /var/run/docker.sock:/tmp/docker.sock \
    -h $HOSTNAME progrium/registrator consul://$PUBLIC_IP:8500
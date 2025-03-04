#!/bin/sh

docker run -d --name sithil_cont \ 
	-e DB_HOST=172.17.0.1 \
	-e DB_NAME=sithil \
	-e DB_USER=root \
	-e DB_PASSWORD=temp123 \
	-e DB_PASSWORD_ROOT=temp123 \
	-e DB_PORT=3306 \
	-e JWT_SECRET=xdd \
	-p 8000:8000 sithil
